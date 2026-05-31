package password

import "golang.org/x/crypto/bcrypt"

func Hash(
    value string,
) (string,error) {

    bytes, err := bcrypt.GenerateFromPassword(
        []byte(value),
        bcrypt.DefaultCost,
    )

    return string(bytes), err
}

func Compare(
    hash string,
    value string,
) bool {

    return bcrypt.CompareHashAndPassword(
        []byte(hash),
        []byte(value),
    ) == nil
}
