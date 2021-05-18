export FIRESTORE_EMULATOR_HOST=localhost:8080
export GCLOUD_PROJECT="mecab_converter"
export FIREBASE_AUTH_EMULATOR_HOST="localhost:9099"
firebase emulators:start --only firestore &

go test ./...
