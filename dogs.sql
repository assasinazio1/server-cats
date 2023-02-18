CREATE TABLE "dogs" (
                        "id" bigserial PRIMARY KEY,
                        "name" varchar NOT NULL,
                        "ear_length" integer NOT NULL DEFAULT 5,
                        "color" varchar