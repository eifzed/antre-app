CREATE TABLE "ant_mst_user" (
	"user_id" bigserial not null,
	"name" varchar(255),
	"phone_number" varchar,
	"email" varchar(255) NOT NULL,
	"address" VARCHAR(255),
	"user_type" varchar NOT NULL,
	"postal_code" integer NOT NULL,
	"create_time" timestamptz NOT NULL DEFAULT now(),
	"update_time" timestamptz NOT NULL DEFAULT now(),
	"delete_time" timestamptz,
	CONSTRAINT "ant_mst_user_pk" PRIMARY KEY ("user_id")
);



CREATE TABLE "ant_dtl_company" (
	"company_id" bigserial not null,
	"company_name" varchar(255) NOT NULL,
	"owner_id" int8 NOT NULL,
	"address" varchar(255) NOT NULL,
	"postal_code" integer NOT NULL,
	"create_time" timestamptz NOT NULL DEFAULT now(),
	"update_time" timestamptz NOT NULL DEFAULT now(),
	"delete_time" timestamptz,
	CONSTRAINT "ant_dtl_company_pk" PRIMARY KEY ("company_id")
);



CREATE TABLE "ant_trx_reservation" (
	"reservation_id" bigserial not null NOT NULL,
	"customer_id" int8 NOT NULL,
	"shop_id" int8 NOT NULL,
	"reservation_time" timestamptz NOT NULL,
	"reservation_type" varchar NOT NULL,
	"status_id" int NOT NULL,
	"customer_note" varchar(255) NOT NULL,
	"create_time" timestamptz NOT NULL DEFAULT now(),
	"update_time" timestamptz NOT NULL DEFAULT now(),
	"delete_time" timestamptz,
	CONSTRAINT "ant_trx_reservation_pk" PRIMARY KEY ("reservation_id")
);



CREATE TABLE "ant_dtl_shop" (
	"shop_id" bigserial NOT NULL,
	"company_id" integer NOT NULL,
	"pic_id" integer NOT NULL,
	"shop_name" varchar(255) NOT NULL,
	"shop_type" varchar NOT NULL,
	"address" varchar(255) NOT NULL,
	"postal_code" integer NOT NULL,
	"create_time" timestamptz NOT NULL DEFAULT now(),
	"update_time" timestamptz NOT NULL DEFAULT now(),
	"delete_time" timestamptz,
	CONSTRAINT "ant_dtl_shop_pk" PRIMARY KEY ("shop_id")
);



CREATE TABLE "ant_hst_reservation" (
	"hst_reservation_id" bigserial NOT NULL,
	"reservation_id" int8 NOT NULL,
	"customer_id" int8 NOT NULL,
	"shop_id" int NOT NULL,
	"reservation_time" timestamptz NOT NULL,
	"reservation_type" varchar NOT NULL,
	"status_id" int NOT NULL,
	"customer_note" varchar(255) NOT NULL,
	"updater_id" int8 NOT NULL,
	"reason" varchar(255) NOT NULL,
	"update_type" varchar NOT NULL,
	"create_time" timestamptz NOT NULL DEFAULT now(),
	"update_time" timestamptz NOT NULL DEFAULT now(),
	"delete_time" timestamptz,
	CONSTRAINT "ant_hst_reservation_pk" PRIMARY KEY ("hst_reservation_id")
);



CREATE TABLE "ant_mst_category" (
	"category_id" bigserial not null,
	"category_level" int NOT NULL,
	"category_name" varchar NOT NULL,
	"create_time" timestamptz NOT NULL DEFAULT now(),
	"update_time" timestamptz NOT NULL DEFAULT now(),
	"delete_time" timestamptz,
	CONSTRAINT "ant_mst_category_pk" PRIMARY KEY ("category_id")
);



CREATE TABLE "ant_map_shop_category" (
	"map_id" bigserial not null,
	"shop_id" int NOT NULL,
	"category_id" int8 NOT NULL,
	"create_time" timestamptz NOT NULL DEFAULT now(),
	"update_time" timestamptz NOT NULL DEFAULT now(),
	"delete_time" timestamptz,
	CONSTRAINT "ant_map_shop_category_pk" PRIMARY KEY ("map_id")
);



CREATE TABLE "ant_mst_status" (
	"status_id" bigserial not null,
	"status_name" varchar NOT NULL,
	"checkpoint_id" int NOT NULL,
	"checkpoint_name" varchar NOT NULL,
	"create_time" timestamptz NOT NULL DEFAULT now(),
	"update_time" timestamptz NOT NULL DEFAULT now(),
	"delete_time" timestamptz,
	CONSTRAINT "ant_mst_status_pk" PRIMARY KEY ("status_id")
);





ALTER TABLE "ant_trx_reservation" ADD CONSTRAINT "ant_trx_reservation_fk0" FOREIGN KEY ("customer_id") REFERENCES "ant_mst_user"("user_id");
ALTER TABLE "ant_trx_reservation" ADD CONSTRAINT "ant_trx_reservation_fk1" FOREIGN KEY ("shop_id") REFERENCES "ant_dtl_shop"("shop_id");
ALTER TABLE "ant_trx_reservation" ADD CONSTRAINT "ant_trx_reservation_fk2" FOREIGN KEY ("status_id") REFERENCES "ant_mst_status"("status_id");

ALTER TABLE "ant_dtl_shop" ADD CONSTRAINT "ant_dtl_shop_fk0" FOREIGN KEY ("company_id") REFERENCES "ant_dtl_company"("company_id");
ALTER TABLE "ant_dtl_shop" ADD CONSTRAINT "ant_dtl_shop_fk1" FOREIGN KEY ("pic_id") REFERENCES "ant_mst_user"("user_id");

ALTER TABLE "ant_hst_reservation" ADD CONSTRAINT "ant_hst_reservation_fk0" FOREIGN KEY ("resrvation_id") REFERENCES "ant_trx_reservation"("reservation_id");
ALTER TABLE "ant_hst_reservation" ADD CONSTRAINT "ant_hst_reservation_fk1" FOREIGN KEY ("customer_id") REFERENCES "ant_mst_user"("user_id");
ALTER TABLE "ant_hst_reservation" ADD CONSTRAINT "ant_hst_reservation_fk2" FOREIGN KEY ("shop_id") REFERENCES "ant_dtl_shop"("shop_id");
ALTER TABLE "ant_hst_reservation" ADD CONSTRAINT "ant_hst_reservation_fk3" FOREIGN KEY ("status_id") REFERENCES "ant_mst_status"("status_id");
ALTER TABLE "ant_hst_reservation" ADD CONSTRAINT "ant_hst_reservation_fk4" FOREIGN KEY ("updater_id") REFERENCES "ant_mst_user"("user_id");


ALTER TABLE "ant_map_shop_category" ADD CONSTRAINT "ant_map_shop_category_fk0" FOREIGN KEY ("shop_id") REFERENCES "ant_dtl_shop"("shop_id");
ALTER TABLE "ant_map_shop_category" ADD CONSTRAINT "ant_map_shop_category_fk1" FOREIGN KEY ("category_id") REFERENCES "ant_mst_category"("category_id");


CREATE INDEX CONCURRENTLY ant_mst_user_user_type_delete_time ON ant_mst_user USING btree (user_type, delete_time);

CREATE INDEX CONCURRENTLY ant_dtl_company_owner_id_delete_time ON ant_dtl_company USING btree (delete_time, owner_id);

CREATE INDEX CONCURRENTLY ant_trx_reservation_customer_id_shop_id_reservation_time_status_id_delete_time ON ant_trx_reservation USING btree (delete_time, customer_id, shop_id, reservation_time, status_id);

CREATE INDEX CONCURRENTLY ant_dtl_shop_pic_id_company_id_delete_time ON ant_dtl_shop USING btree (delete_time, pic_id, company_id);

CREATE INDEX CONCURRENTLY ant_hst_reservation_reservation_id_customer_id_shop_id_delete_time ON ant_hst_reservation USING btree (reservation_id, customer_id, shop_id, delete_time);

CREATE INDEX CONCURRENTLY ant_map_shop_category_shop_id_category_id_delete_time ON ant_map_shop_category USING btree (shop_id, category_id, delete_time);





