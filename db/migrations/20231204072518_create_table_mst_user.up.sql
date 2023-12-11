CREATE TABLE "mst_user" (
                                "id" uuid DEFAULT uuid_generate_v4(),
                                "username" varchar(255),
                                "email" varchar(255),
                                "password" varchar(255),
                                "created_date" date NOT NULL DEFAULT CURRENT_DATE,
                                "updated_date" date NOT NULL DEFAULT CURRENT_DATE,
                                PRIMARY KEY("id")
);
