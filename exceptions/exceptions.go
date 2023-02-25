package exceptions

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/upper-institute/go-kit/logging"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func MustMarshalJson(src map[string]interface{}) string {

	raw, err := json.Marshal(src)
	if err != nil {
		panic(err)
	}

	return string(raw)

}

func ThrowInternalErr(err error) error {

	if _, ok := status.FromError(err); ok {
		return err
	}

	errHash := md5.New()
	errHash.Write([]byte(err.Error()))
	hash := hex.EncodeToString(errHash.Sum(nil))

	log := logging.Logger.Named("exceptions").Sugar()

	log.Errorw("Internal error", "hash", hash, "error", err)

	return status.Error(codes.Internal, fmt.Sprintf(`{"internalError":"%s"}`, hash))

}
