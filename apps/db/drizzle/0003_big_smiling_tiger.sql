ALTER TABLE "accounts" ADD COLUMN "balance" numeric(12, 2) DEFAULT 0 NOT NULL;--> statement-breakpoint
ALTER TABLE "accounts" ADD COLUMN "currency" varchar(3) DEFAULT 'USD' NOT NULL;