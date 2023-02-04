CREATE TABLE "cats" (
                        "id" bigserial PRIMARY KEY,
                        "name" varchar NOT NULL,
                        "is_strip" boolean NOT NULL DEFAULT false,
                        "color" varchar
)