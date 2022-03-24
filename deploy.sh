#! /bin/bash
# 4 types of triggers:
create_trigger="providers/cloud.firestore/eventTypes/document.create"
update_trigger="providers/cloud.firestore/eventTypes/document.update"
delete_trigger="providers/cloud.firestore/eventTypes/document.delete"
write_trigger="providers/cloud.firestore/eventTypes/document.write"

projectid="lcwp-jack"
if [ "$(git branch --show-current)" = "main" ]; then
  echo "@ PROD"
  runmode="prod"
elif [ "$(git branch --show-current)" = "dev" ]; then
  echo "@ DEV"
  runmode="dev"
else
  echo "@ You're not in the main or dev branch , so....byebye~~"
  exit 0
fi
set -x
gcloud functions deploy ${runmode}-hello-world \
  --project="${projectid}" \
  --region=asia-east1\
  --service-account firebase-adminsdk-dyqdv@lcwp-jack.iam.gserviceaccount.com \
  --entry-point EntryHelloWorld \
  --runtime go116 \
  --trigger-event ${create_trigger} \
  --trigger-resource "projects/${projectid}/databases/(default)/documents/Example/{eid}" \
  --allow-unauthenticated \
  --memory 128

gcloud functions deploy ${runmode}-bye-world \
  --project="${projectid}" \
  --region=asia-east1\
  --service-account firebase-adminsdk-dyqdv@lcwp-jack.iam.gserviceaccount.com \
  --entry-point EntryByeWorld \
  --runtime go116 \
  --trigger-event ${delete_trigger} \
  --trigger-resource "projects/${projectid}/databases/(default)/documents/Example/{eid}" \
  --allow-unauthenticated \
  --memory 128