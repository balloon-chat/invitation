package firebase

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go/v4"
	"fmt"
	"github.com/balloon/go/invite/internal/domain/model"
	"github.com/balloon/go/invite/internal/domain/repository"
	"github.com/balloon/go/invite/internal/domain/util"
	"google.golang.org/api/iterator"
	"sync"
)

var client *firestore.Client

const (
	invitationCollectionKey = "invitations"
)

type FirestoreInvitationRepository struct {
	lock       sync.RWMutex
	collection *firestore.CollectionRef
}

func NewFirestoreInvitationRepository() (*FirestoreInvitationRepository, error) {
	if client == nil {
		ctx := context.Background()
		app, err := firebase.NewApp(ctx, nil)
		if err != nil {
			return nil, fmt.Errorf("error while initializing firebase app: %v", err)
		}
		client, err = app.Firestore(ctx)
		if err != nil {
			return nil, fmt.Errorf("error while initializing firestore: %v", err)
		}
	}

	r := &FirestoreInvitationRepository{
		collection: client.Collection(invitationCollectionKey),
	}

	return r, nil
}

func (r *FirestoreInvitationRepository) Save(invitation *model.Invitation) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	entity := repository.NewInvitationEntity(*invitation)
	_, _, err := r.collection.Add(context.Background(), entity.ToMap())
	if err != nil {
		return fmt.Errorf("error while writing to firestore: %v", err)
	}
	return nil
}

func (r *FirestoreInvitationRepository) FindByTopicId(topicId model.TopicId) (*model.Invitation, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	docsItr := r.collection.Where(repository.KeyTopicId, "==", topicId).Documents(context.Background())
	doc, err := docsItr.Next()
	if err == iterator.Done {
		return nil, repository.InvitationNotFoundError
	}
	if err != nil {
		return nil, fmt.Errorf("error while getting data from firestore: %v", err)
	}

	var entity repository.InvitationEntity
	err = doc.DataTo(&entity)
	if err != nil {
		return nil, fmt.Errorf("error while converting data: %v", entity)
	}

	return entity.ToInvitation(), nil
}

func (r *FirestoreInvitationRepository) FindByInvitationCode(code model.InvitationCode) (*model.Invitation, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	docItr := r.collection.Where(repository.KeyCode, "==", util.ArrayToInt(code)).Documents(context.Background())
	doc, err := docItr.Next()
	if err == iterator.Done {
		return nil, repository.InvitationNotFoundError
	}
	if err != nil {
		return nil, fmt.Errorf("error while getting data from firestore: %v", err)
	}

	var entity repository.InvitationEntity
	err = doc.DataTo(&entity)
	if err != nil {
		return nil, fmt.Errorf("error while converting data: %v", err)
	}

	return entity.ToInvitation(), nil
}
