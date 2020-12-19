package address

import (
	"errors"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"github.com/thomasxnguy/bitcoinaddress/common"
	"github.com/thomasxnguy/bitcoinaddress/database"
	apierrors "github.com/thomasxnguy/bitcoinaddress/errors"
	"github.com/thomasxnguy/bitcoinaddress/models"
	"net/http"
	"time"
)

// Service to generate and manage bitcoin addresses
type Service struct {
	AccountStore database.AccountStorer
	KeyManager   *common.KeyManager
	counter      common.Count
}

// NewService create a new address service
func NewService(accountStore database.AccountStorer) *Service {
	mnemonic := viper.GetString("mnemonic")
	password := viper.GetString("password")
	testnet := viper.GetBool("testnet")
	return &Service{
		AccountStore: accountStore,
		KeyManager:   common.NewKeyManager(mnemonic, password, testnet),
	}
}

// Endpoint to generate addresses for user. A new account identified by an UUID will be created and mapped to those addresses.
func (rs *Service) generateAddresses(w http.ResponseWriter, r *http.Request) {
	// Generate a new id.
	userId := uuid.New()

	// Create a new account, we use the atomic count to generate the key index.
	var newUser = models.Account{
		Id:        userId,
		KeyIndex:  rs.counter.Inc(),
		CreatedAt: time.Now(),
	}
	rs.AccountStore.Create(&newUser)

	// Get Segwit Address of user
	segwitAddress, err := rs.KeyManager.GetSegWitAddressForAccountAt(newUser.KeyIndex)
	if err != nil {
		render.Render(w, r, apierrors.ErrInternalError(err))
		return
	}
	// Get Native Segwit Address of user
	nativeSegwitAddress, err := rs.KeyManager.GetNativeSegWitAddressForAccountAt(newUser.KeyIndex)
	if err != nil {
		render.Render(w, r, apierrors.ErrInternalError(err))
		return
	}

	render.Respond(w, r, newGenerateAddressResponse(&userId, segwitAddress, nativeSegwitAddress))
}

// Endpoint to get the addresses of a particular user. Return not found is user does not exists.
func (rs *Service) getUserAddresses(w http.ResponseWriter, r *http.Request) {
	userIdParam := chi.URLParam(r, "user_id")
	userId, err := uuid.Parse(userIdParam)
	if err != nil {
		render.Render(w, r, apierrors.ErrInvalidRequest(errors.New("Must be a valid user_id")))
		return
	}

	account, _ := rs.AccountStore.Get(userId)
	if account == nil {
		render.Render(w, r, apierrors.ErrNotFound(errors.New("User is not found")))
		return
	}

	// Get Segwit Address of user
	segwitAddress, err := rs.KeyManager.GetSegWitAddressForAccountAt(account.KeyIndex)
	if err != nil {
		render.Render(w, r, apierrors.ErrInternalError(err))
		return
	}
	// Get Native Segwit Address of user
	nativeSegwitAddress, err := rs.KeyManager.GetNativeSegWitAddressForAccountAt(account.KeyIndex)
	if err != nil {
		render.Render(w, r, apierrors.ErrInternalError(err))
		return
	}

	render.Respond(w, r, newGetAddressResponse(segwitAddress, nativeSegwitAddress))
}
