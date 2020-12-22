FROM amacneil/dbmate

COPY ./.env /
ADD ./db-chckr /db-chckr
ADD ./db-checkins /db-checkins

RUN apk update && apk add bash
ENTRYPOINT ["/bin/bash"]
