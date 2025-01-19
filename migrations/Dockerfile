FROM migrate/migrate

ENV DB_URI=uri

WORKDIR /migrations

COPY . .

ENTRYPOINT migrate -database $DB_URI -path ./ up
