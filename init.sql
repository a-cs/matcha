CREATE TABLE "blocks"(
    "id" SERIAL PRIMARY KEY,
    "blocker_user_id" INTEGER NOT NULL,
    "blecked_user_id" INTEGER NOT NULL,
    "is_active" BOOLEAN NOT NULL
);
CREATE TABLE "messages"(
    "id" SERIAL PRIMARY KEY,
    "match_id" INTEGER NOT NULL,
    "sender_user_id" INTEGER NOT NULL,
    "content" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
CREATE TABLE "notifications_types"(
    "id" SERIAL PRIMARY KEY,
    "type" VARCHAR(255) NOT NULL
);
CREATE TABLE "profile"(
    "id" SERIAL PRIMARY KEY,
    "user_id" INTEGER UNIQUE NOT NULL,
    "first_name" VARCHAR(255) NOT NULL,
    "last_name" VARCHAR(255) NOT NULL,
    "location" VARCHAR(255),
    "likes_counter" INTEGER NOT NULL,
    "gender_id" INTEGER NOT NULL,
    "tags_list" JSON NOT NULL,
    "biography" VARCHAR(255),
    "sexual_preference_id" INTEGER NOT NULL,
    "pictures" JSON,
    "view_counter" INTEGER NOT NULL,
    "is_online" BOOLEAN NOT NULL,
    "last_online_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
    "account_status" VARCHAR(255) NOT NULL
);
CREATE TABLE "likes_log"(
    "id" SERIAL PRIMARY KEY,
    "giver_user_id" INTEGER NOT NULL,
    "receiver_user_id" INTEGER NOT NULL,
    "is_active" BOOLEAN NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "last_updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
CREATE TABLE "users"(
    "id" SERIAL PRIMARY KEY,
    "email" VARCHAR(255) UNIQUE NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "username" VARCHAR(255) UNIQUE NOT NULL,
    "active_matches" JSON NOT NULL,
    "account_status" VARCHAR(255) NOT NULL
    "slug_id" VARCHAR(255) NOT NULL
);
CREATE TABLE "sexual_preference"(
    "id" SERIAL PRIMARY KEY,
    "option" VARCHAR(255) NOT NULL
);
CREATE TABLE "notification"(
    "id" SERIAL PRIMARY KEY,
    "to_user_id" INTEGER NOT NULL,
    "from_user_id" INTEGER NOT NULL,
    "content" VARCHAR(255) NOT NULL,
    "type_id" INTEGER NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "was_read" BOOLEAN NOT NULL
);
CREATE TABLE "gender"(
    "id" SERIAL PRIMARY KEY,
    "type" VARCHAR(255) NOT NULL
);
CREATE TABLE "match"(
    "id" SERIAL PRIMARY KEY,
    "giver_user_id" INTEGER NOT NULL,
    "receiver_user_id" INTEGER NOT NULL,
    "status" VARCHAR(255) NOT NULL,
    "is_active" BOOLEAN NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "last_updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
CREATE TABLE "tags"(
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL
);
ALTER TABLE
    "profile" ADD CONSTRAINT "profile_gender_id_foreign" FOREIGN KEY("gender_id") REFERENCES "gender"("id");
ALTER TABLE
    "notification" ADD CONSTRAINT "notification_type_id_foreign" FOREIGN KEY("type_id") REFERENCES "notifications_types"("id");
ALTER TABLE
    "messages" ADD CONSTRAINT "messages_match_id_foreign" FOREIGN KEY("match_id") REFERENCES "match"("id");
ALTER TABLE
    "profile" ADD CONSTRAINT "profile_sexual_preference_id_foreign" FOREIGN KEY("sexual_preference_id") REFERENCES "sexual_preference"("id");
ALTER TABLE
    "profile" ADD CONSTRAINT "profile_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("id");