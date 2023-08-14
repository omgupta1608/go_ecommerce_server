CREATE TABLE users
(
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    name text COLLATE pg_catalog."default" NOT NULL,
    email text COLLATE pg_catalog."default" NOT NULL,
    password text COLLATE pg_catalog."default" NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone,
    updated_at timestamp with time zone,
    CONSTRAINT users_pkey PRIMARY KEY (id)
);

CREATE TABLE ref_order_status
(
    id text COLLATE pg_catalog."default" NOT NULL,
    description text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT ref_order_status_pkey PRIMARY KEY (id)
);

CREATE TABLE ratings
(
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    product_id uuid NOT NULL,
    user_id uuid NOT NULL,
    rating integer NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    CONSTRAINT ratings_pkey PRIMARY KEY (id),
    CONSTRAINT ratings_user_fkey FOREIGN KEY (user_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE RESTRICT
        ON DELETE RESTRICT
);

CREATE TABLE products
(
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    name text COLLATE pg_catalog."default" NOT NULL,
    price double precision NOT NULL,
    in_stock integer NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    CONSTRAINT products_pkey PRIMARY KEY (id)
);


CREATE TABLE orders
(
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    user_id uuid NOT NULL,
    status text COLLATE pg_catalog."default" NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    CONSTRAINT orders_pkey PRIMARY KEY (id),
    CONSTRAINT order_status_fkey FOREIGN KEY (status)
        REFERENCES public.ref_order_status (id) MATCH SIMPLE
        ON UPDATE RESTRICT
        ON DELETE RESTRICT,
    CONSTRAINT order_user_fkey FOREIGN KEY (user_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE RESTRICT
        ON DELETE RESTRICT
);


CREATE TABLE order_products
(
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    order_id uuid NOT NULL,
    product_id uuid NOT NULL,
    quantity integer NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    placed boolean NOT NULL DEFAULT true,
    CONSTRAINT order_products_pkey PRIMARY KEY (id),
    CONSTRAINT order_products_order_fkey FOREIGN KEY (order_id)
        REFERENCES public.orders (id) MATCH SIMPLE
        ON UPDATE RESTRICT
        ON DELETE RESTRICT,
    CONSTRAINT order_products_product_fkey FOREIGN KEY (product_id)
        REFERENCES public.products (id) MATCH SIMPLE
        ON UPDATE RESTRICT
        ON DELETE RESTRICT
);


alter table "public"."orders"
  add constraint "order_status_fkey"
  foreign key ("status")
  references "public"."ref_order_status"
  ("id") on update restrict on delete restrict;

alter table "public"."orders"
  add constraint "order_user_fkey"
  foreign key ("user_id")
  references "public"."users"
  ("id") on update restrict on delete restrict;

alter table "public"."order_products"
  add constraint "order_products_product_fkey"
  foreign key ("product_id")
  references "public"."products"
  ("id") on update restrict on delete restrict;

alter table "public"."order_products"
  add constraint "order_products_order_fkey"
  foreign key ("order_id")
  references "public"."orders"
  ("id") on update restrict on delete restrict;

alter table "public"."ratings"
  add constraint "ratings_user_fkey"
  foreign key ("user_id")
  references "public"."users"
  ("id") on update restrict on delete restrict;

