CREATE TABLE "categories" (
	"id" varchar PRIMARY KEY NOT NULL,
	"name" varchar NOT NULL
);
--> statement-breakpoint
ALTER TABLE "accounts" ADD COLUMN "account_number" varchar;--> statement-breakpoint
ALTER TABLE "transactions" ADD COLUMN "categoryId" varchar;--> statement-breakpoint
ALTER TABLE "transactions" ADD CONSTRAINT "transactions_categoryId_categories_id_fk" FOREIGN KEY ("categoryId") REFERENCES "public"."categories"("id") ON DELETE no action ON UPDATE no action;