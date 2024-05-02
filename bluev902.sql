--
-- PostgreSQL database dump
--

-- Dumped from database version 15.6 (Ubuntu 15.6-0ubuntu0.23.10.1)
-- Dumped by pg_dump version 16.2 (Ubuntu 16.2-1.pgdg23.10+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: apps; Type: TABLE; Schema: public; Owner: blueuser
--

CREATE TABLE public.apps (
    id bigint NOT NULL,
    name text,
    uuid uuid,
    active boolean DEFAULT true,
    description text
);


ALTER TABLE public.apps OWNER TO blueuser;

--
-- Name: apps_id_seq; Type: SEQUENCE; Schema: public; Owner: blueuser
--

CREATE SEQUENCE public.apps_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.apps_id_seq OWNER TO blueuser;

--
-- Name: apps_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: blueuser
--

ALTER SEQUENCE public.apps_id_seq OWNED BY public.apps.id;


--
-- Name: blob_pictures; Type: TABLE; Schema: public; Owner: blueuser
--

CREATE TABLE public.blob_pictures (
    id bigint NOT NULL,
    blob_picture bytea
);


ALTER TABLE public.blob_pictures OWNER TO blueuser;

--
-- Name: blob_pictures_id_seq; Type: SEQUENCE; Schema: public; Owner: blueuser
--

CREATE SEQUENCE public.blob_pictures_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.blob_pictures_id_seq OWNER TO blueuser;

--
-- Name: blob_pictures_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: blueuser
--

ALTER SEQUENCE public.blob_pictures_id_seq OWNED BY public.blob_pictures.id;


--
-- Name: blob_videos; Type: TABLE; Schema: public; Owner: blueuser
--

CREATE TABLE public.blob_videos (
    id bigint NOT NULL,
    name text,
    blob_video bytea
);


ALTER TABLE public.blob_videos OWNER TO blueuser;

--
-- Name: blob_videos_id_seq; Type: SEQUENCE; Schema: public; Owner: blueuser
--

CREATE SEQUENCE public.blob_videos_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.blob_videos_id_seq OWNER TO blueuser;

--
-- Name: blob_videos_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: blueuser
--

ALTER SEQUENCE public.blob_videos_id_seq OWNED BY public.blob_videos.id;


--
-- Name: end_points; Type: TABLE; Schema: public; Owner: blueuser
--

CREATE TABLE public.end_points (
    id bigint NOT NULL,
    name text,
    route_paths text,
    method text,
    description text,
    feature_id bigint
);


ALTER TABLE public.end_points OWNER TO blueuser;

--
-- Name: end_points_id_seq; Type: SEQUENCE; Schema: public; Owner: blueuser
--

CREATE SEQUENCE public.end_points_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.end_points_id_seq OWNER TO blueuser;

--
-- Name: end_points_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: blueuser
--

ALTER SEQUENCE public.end_points_id_seq OWNED BY public.end_points.id;


--
-- Name: features; Type: TABLE; Schema: public; Owner: blueuser
--

CREATE TABLE public.features (
    id bigint NOT NULL,
    name text,
    description text,
    active boolean DEFAULT true,
    role_id bigint
);


ALTER TABLE public.features OWNER TO blueuser;

--
-- Name: features_id_seq; Type: SEQUENCE; Schema: public; Owner: blueuser
--

CREATE SEQUENCE public.features_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.features_id_seq OWNER TO blueuser;

--
-- Name: features_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: blueuser
--

ALTER SEQUENCE public.features_id_seq OWNED BY public.features.id;


--
-- Name: jwt_salts; Type: TABLE; Schema: public; Owner: blueuser
--

CREATE TABLE public.jwt_salts (
    id bigint NOT NULL,
    salt_a text,
    salt_b text
);


ALTER TABLE public.jwt_salts OWNER TO blueuser;

--
-- Name: jwt_salts_id_seq; Type: SEQUENCE; Schema: public; Owner: blueuser
--

CREATE SEQUENCE public.jwt_salts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.jwt_salts_id_seq OWNER TO blueuser;

--
-- Name: jwt_salts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: blueuser
--

ALTER SEQUENCE public.jwt_salts_id_seq OWNED BY public.jwt_salts.id;


--
-- Name: page_roles; Type: TABLE; Schema: public; Owner: blueuser
--

CREATE TABLE public.page_roles (
    page_id bigint NOT NULL,
    role_id bigint NOT NULL
);


ALTER TABLE public.page_roles OWNER TO blueuser;

--
-- Name: pages; Type: TABLE; Schema: public; Owner: blueuser
--

CREATE TABLE public.pages (
    id bigint NOT NULL,
    name text,
    active boolean DEFAULT true,
    description text
);


ALTER TABLE public.pages OWNER TO blueuser;

--
-- Name: pages_id_seq; Type: SEQUENCE; Schema: public; Owner: blueuser
--

CREATE SEQUENCE public.pages_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.pages_id_seq OWNER TO blueuser;

--
-- Name: pages_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: blueuser
--

ALTER SEQUENCE public.pages_id_seq OWNED BY public.pages.id;


--
-- Name: roles; Type: TABLE; Schema: public; Owner: blueuser
--

CREATE TABLE public.roles (
    id bigint NOT NULL,
    name text NOT NULL,
    description text NOT NULL,
    active boolean DEFAULT true,
    app_id bigint
);


ALTER TABLE public.roles OWNER TO blueuser;

--
-- Name: roles_id_seq; Type: SEQUENCE; Schema: public; Owner: blueuser
--

CREATE SEQUENCE public.roles_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.roles_id_seq OWNER TO blueuser;

--
-- Name: roles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: blueuser
--

ALTER SEQUENCE public.roles_id_seq OWNED BY public.roles.id;


--
-- Name: session_data; Type: TABLE; Schema: public; Owner: blueuser
--

CREATE TABLE public.session_data (
    token text,
    time_stamp timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.session_data OWNER TO blueuser;

--
-- Name: site_data; Type: TABLE; Schema: public; Owner: blueuser
--

CREATE TABLE public.site_data (
    id bigint NOT NULL,
    remote_add character varying(128),
    accessed_route character varying(300),
    method character varying(10),
    response_time numeric,
    response_status bigint
);


ALTER TABLE public.site_data OWNER TO blueuser;

--
-- Name: site_data_id_seq; Type: SEQUENCE; Schema: public; Owner: blueuser
--

CREATE SEQUENCE public.site_data_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.site_data_id_seq OWNER TO blueuser;

--
-- Name: site_data_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: blueuser
--

ALTER SEQUENCE public.site_data_id_seq OWNED BY public.site_data.id;


--
-- Name: user_roles; Type: TABLE; Schema: public; Owner: blueuser
--

CREATE TABLE public.user_roles (
    role_id bigint NOT NULL,
    user_id bigint NOT NULL
);


ALTER TABLE public.user_roles OWNER TO blueuser;

--
-- Name: users; Type: TABLE; Schema: public; Owner: blueuser
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    uuid uuid,
    email text,
    password text,
    date_registered timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    disabled boolean DEFAULT false
);


ALTER TABLE public.users OWNER TO blueuser;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: blueuser
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO blueuser;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: blueuser
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: apps id; Type: DEFAULT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.apps ALTER COLUMN id SET DEFAULT nextval('public.apps_id_seq'::regclass);


--
-- Name: blob_pictures id; Type: DEFAULT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.blob_pictures ALTER COLUMN id SET DEFAULT nextval('public.blob_pictures_id_seq'::regclass);


--
-- Name: blob_videos id; Type: DEFAULT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.blob_videos ALTER COLUMN id SET DEFAULT nextval('public.blob_videos_id_seq'::regclass);


--
-- Name: end_points id; Type: DEFAULT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.end_points ALTER COLUMN id SET DEFAULT nextval('public.end_points_id_seq'::regclass);


--
-- Name: features id; Type: DEFAULT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.features ALTER COLUMN id SET DEFAULT nextval('public.features_id_seq'::regclass);


--
-- Name: jwt_salts id; Type: DEFAULT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.jwt_salts ALTER COLUMN id SET DEFAULT nextval('public.jwt_salts_id_seq'::regclass);


--
-- Name: pages id; Type: DEFAULT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.pages ALTER COLUMN id SET DEFAULT nextval('public.pages_id_seq'::regclass);


--
-- Name: roles id; Type: DEFAULT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.roles ALTER COLUMN id SET DEFAULT nextval('public.roles_id_seq'::regclass);


--
-- Name: site_data id; Type: DEFAULT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.site_data ALTER COLUMN id SET DEFAULT nextval('public.site_data_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: apps; Type: TABLE DATA; Schema: public; Owner: blueuser
--

COPY public.apps (id, name, uuid, active, description) FROM stdin;
1	BlueAdmin	48015a9b-5a86-4a15-944b-94108aa78b4b	t	SSO Role Based User Administration solution
2	BlueCom	21028fa1-8e04-464e-8be7-95ea9c82994b	t	Commercial Order Management
3	BlueHRM	9359b1ba-98b5-427c-96d8-023fc33cd1b0	t	Human Resource Management
\.


--
-- Data for Name: blob_pictures; Type: TABLE DATA; Schema: public; Owner: blueuser
--

COPY public.blob_pictures (id, blob_picture) FROM stdin;
\.


--
-- Data for Name: blob_videos; Type: TABLE DATA; Schema: public; Owner: blueuser
--

COPY public.blob_videos (id, name, blob_video) FROM stdin;
\.


--
-- Data for Name: end_points; Type: TABLE DATA; Schema: public; Owner: blueuser
--

COPY public.end_points (id, name, route_paths, method, description, feature_id) FROM stdin;
1	swagger_routes_get	/docs/*	GET	swagger_routes-GET	\N
2	custom_metrics_route_get	/lmetrics	GET	custom_metrics_route-GET	13
3	check_login_get	/api/v1/checklogin	GET	check_login-GET	13
4	roles_get	/api/v1/roles	GET	roles-GET	1
5	roles_single_get	/api/v1/roles/:id	GET	roles_single-GET	1
6	drop_roles_get	/api/v1/droproles	GET	drop_roles-GET	14
7	roles_endpoints_get	/api/v1/role_endpoints	GET	roles_endpoints-GET	1
8	features_get	/api/v1/features	GET	features-GET	11
9	features_single_get	/api/v1/features/:id	GET	features_single-GET	11
10	drop_features_get	/api/v1/featuredrop	GET	drop_features-GET	14
11	apps_get	/api/v1/apps	GET	apps-GET	7
12	apps_single_get	/api/v1/apps/:id	GET	apps_single-GET	7
13	drop_sppd_get	/api/v1/appsdrop	GET	drop_sppd-GET	14
14	apps_features_get	/api/v1/appsmatrix/:id	GET	apps_features-GET	7
15	users_get	/api/v1/users	GET	users-GET	3
16	user_single_get	/api/v1/users/:id	GET	user_single-GET	3
17	get_user_roles_get	/api/v1/userrole/:user_id	GET	get_user_roles-GET	3
18	pages_get	/api/v1/pages	GET	pages-GET	5
19	page_single_get	/api/v1/pages/:id	GET	page_single-GET	5
20	get_page_roles_get	/api/v1/pagesroles/:page_id	GET	get_page_roles-GET	5
21	end_point_get	/api/v1/endpoints	GET	end_point-GET	9
22	end_point_single_get	/api/v1/endpoints/:id	GET	end_point_single-GET	9
23	drop_endpoints_get	/api/v1/endpointdrop	GET	drop_endpoints-GET	14
25	login_route_post	/api/v1/login	POST	login_route-POST	13
26	roles_post	/api/v1/roles	POST	roles-POST	2
27	features_post	/api/v1/features	POST	features-POST	12
28	apps_post	/api/v1/apps	POST	apps-POST	8
29	users_post	/api/v1/users	POST	users-POST	4
30	user_role_post	/api/v1/userrole/:user_id/:role_id	POST	user_role-POST	4
31	pages_post	/api/v1/pages	POST	pages-POST	6
32	page_roles_post	/api/v1/pagerole/:page_id/:role_id	POST	page_roles-POST	6
33	end_point_post	/api/v1/endpoints	POST	end_point-POST	10
34	send_email_post	/api/v1/email	POST	send_email-POST	\N
35	blob_picture_post	/api/v1/blobpic	POST	blob_picture-POST	13
36	blob_video_post	/api/v1/blobvideo	POST	blob_video-POST	13
37	activate_deactivate_role_put	/api/v1/roles/:role_id	PUT	activate_deactivate_role-PUT	2
38	activate_deactivate_features_put	/api/v1/features/:feature_id	PUT	activate_deactivate_features-PUT	12
39	activate_deactivate_user_put	/api/v1/users/:user_id	PUT	activate_deactivate_user-PUT	4
40	roles_single_delete	/api/v1/roles/:id	DELETE	roles_single-DELETE	2
41	features_single_delete	/api/v1/features/:id	DELETE	features_single-DELETE	12
42	feature_role_delete	/api/v1/featuresrole/:feature_id	DELETE	feature_role-DELETE	12
43	apps_single_delete	/api/v1/apps/:id	DELETE	apps_single-DELETE	8
44	user_single_delete	/api/v1/users/:id	DELETE	user_single-DELETE	4
45	user_role_delete	/api/v1/userrole/:user_id/:role_id	DELETE	user_role-DELETE	4
46	page_single_delete	/api/v1/pages/:id	DELETE	page_single-DELETE	6
47	page_roles_delete	/api/v1/pagerole/:page_id/:role_id	DELETE	page_roles-DELETE	6
48	end_point_single_delete	/api/v1/endpoints/:id	DELETE	end_point_single-DELETE	10
49	feature_endpoint_delete	/api/v1/feature_endpoint/:endpoint_id	DELETE	feature_endpoint-DELETE	12
50	roles_single_patch	/api/v1/roles/:id	PATCH	roles_single-PATCH	2
51	roles_app_patch	/api/v1/approle/:role_id	PATCH	roles_app-PATCH	2
52	features_single_patch	/api/v1/features/:id	PATCH	features_single-PATCH	12
53	feature_role_patch	/api/v1/featuresrole/:feature_id	PATCH	feature_role-PATCH	12
54	apps_single_patch	/api/v1/apps/:id	PATCH	apps_single-PATCH	8
55	user_single_patch	/api/v1/users/:id	PATCH	user_single-PATCH	4
56	page_single_patch	/api/v1/pages/:id	PATCH	page_single-PATCH	6
57	end_point_single_patch	/api/v1/endpoints/:id	PATCH	end_point_single-PATCH	10
58	feature_endpoint_patch	/api/v1/feature_endpoint/:endpoint_id	PATCH	feature_endpoint-PATCH	12
59	change_reset_password_put	/api/v1/users/:email_id	PUT	change_reset_password-PUT	4
24	dashboard_get	/api/v1/dashboard	GET	dashboard-GET	13
25	dashboard_one_get	/api/v1/dashboard	GET	dashboard_one-GET	13
26	dashboard_two_get	/api/v1/dashboardends	GET	dashboard_two-GET	13
27	dashboard_three_get	/api/v1/dashboardfeat	GET	dashboard_three-GET	13
28	dashboard_four_get	/api/v1/dashboardpages	GET	dashboard_four-GET	13
29	dashboard_five_get	/api/v1/dashboardroles	GET	dashboard_five-GET	13
30	dashboard_six_get	/api/v1/dashboardrolespage	GET	dashboard_six-GET	13
\.


--
-- Data for Name: features; Type: TABLE DATA; Schema: public; Owner: blueuser
--

COPY public.features (id, name, description, active, role_id) FROM stdin;
1	role_read	View List of Roles	t	2
2	role_write	Privilege to Update ,Create, View and Delete  	t	3
3	user_read	View List of Users	t	2
4	user_write	Privilege to Update ,Create, View and Delete  	t	3
5	page_read	View List of Pages	t	2
6	page_write	Privilege to Update ,Create, View and Delete  	t	3
7	app_read	View List of Apps	t	2
8	app_write	Privilege to Update ,Create, View and Delete 	t	4
9	endpoint_read	View List of Endpoints	t	2
10	endpoint_write	Privilege to Update ,Create, View and Delete  	t	3
11	feature_read	View List of Features	t	2
12	features_write	Privilege to Update ,Create, View and Delete  	t	3
13	login	Endpoints that can be accessed with out Logging In	t	5
14	drop_down	Endpoints that fetch Drowns 	t	6
\.


--
-- Data for Name: jwt_salts; Type: TABLE DATA; Schema: public; Owner: blueuser
--

COPY public.jwt_salts (id, salt_a, salt_b) FROM stdin;
1	1ZkexW3kX7bJDFYacK3EMiHmH	kn5oH9YXNpmg8ikMF19UuRxJm
\.


--
-- Data for Name: page_roles; Type: TABLE DATA; Schema: public; Owner: blueuser
--

COPY public.page_roles (page_id, role_id) FROM stdin;
1	5
5	3
2	5
3	3
3	1
2	1
1	1
4	2
4	3
4	1
5	2
5	1
6	1
6	2
6	3
7	1
7	3
8	1
8	3
7	2
3	2
8	2
2	2
9	1
\.


--
-- Data for Name: pages; Type: TABLE DATA; Schema: public; Owner: blueuser
--

COPY public.pages (id, name, active, description) FROM stdin;
1	Login	t	Login
2	Home	t	Home
3	Role	t	View Roles
4	User	t	View Users
5	Page	t	View Pages
6	Feature	t	View Features
7	Endpoint	t	View Endpoints
8	App	t	View Apps
9	Sign Up	t	Sign Up Page For users Who have not registered
\.


--
-- Data for Name: roles; Type: TABLE DATA; Schema: public; Owner: blueuser
--

COPY public.roles (id, name, description, active, app_id) FROM stdin;
1	superuser	Have Access to All resources in the Apps	t	1
2	standard	Have Access to Limited Features	t	1
3	administrator	Have Access System Admin Privileges	t	1
4	app_role	Have App CURD abilities	t	1
5	Anonymous	For Pages that do not need user sign in	t	1
6	Drop Down 	To access Drop Down menu fetching  endpoints	t	1
\.


--
-- Data for Name: session_data; Type: TABLE DATA; Schema: public; Owner: blueuser
--

COPY public.session_data (token, time_stamp) FROM stdin;
\.


--
-- Data for Name: site_data; Type: TABLE DATA; Schema: public; Owner: blueuser
--

COPY public.site_data (id, remote_add, accessed_route, method, response_time, response_status) FROM stdin;
\.


--
-- Data for Name: user_roles; Type: TABLE DATA; Schema: public; Owner: blueuser
--

COPY public.user_roles (role_id, user_id) FROM stdin;
1	1
1	4
1	5
2	2
2	6
2	7
3	3
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: blueuser
--

COPY public.users (id, uuid, email, password, date_registered, disabled) FROM stdin;
4	8a200cd4-9067-4508-b93c-4b242ef03740	beimnet.degefu@gmail.com	089925fe07a4819e2f032fc4e7b6c1e191dc5d59dbf0bd6b7857a87ef1f7e5c6915ce7ca93f74b136417e362752f81ac5258f2f0be890eff3b7f7679e2cc1b7a	2024-03-08 19:49:53.754679+03	f
6	79731de0-de10-4f62-bcc5-7637d132110a	mickyasne123@gmail.com	3d5267ceaa0759b8756837e4e817517fae8fa2d168267053e78428ac0941d5542c7bee61536be21f7d117fed80bc9aa27e2d450e261ab361ece2a701a1e2da72	2024-03-12 13:52:01.052905+03	f
7	586c2bc7-9a66-49df-a81e-93b2efaec8c9	mickyasne12@gmail.com	adfebcc24b8ce4be96b83381f08275bbd6fe355dee6a1a1407247cd15bbb8ec5952b5680af51c532d8785f950e4bb7ddeab72efc737c4656bcd27b7f73b72bf1	2024-03-12 15:03:48.01416+03	f
1	38ca7360-0138-4b0f-8985-b307ad188e92	superuser@mail.com	089925fe07a4819e2f032fc4e7b6c1e191dc5d59dbf0bd6b7857a87ef1f7e5c6915ce7ca93f74b136417e362752f81ac5258f2f0be890eff3b7f7679e2cc1b7a	2024-01-16 11:27:55.628203+03	f
2	1323b2d9-5755-4e4c-9af5-93c17f59e6fd	standarduser@mail.com	089925fe07a4819e2f032fc4e7b6c1e191dc5d59dbf0bd6b7857a87ef1f7e5c6915ce7ca93f74b136417e362752f81ac5258f2f0be890eff3b7f7679e2cc1b7a	2024-01-17 15:33:56.443849+03	f
3	12bc4954-487b-4263-8bdc-cacbf720f623	adminuser@mail.com	089925fe07a4819e2f032fc4e7b6c1e191dc5d59dbf0bd6b7857a87ef1f7e5c6915ce7ca93f74b136417e362752f81ac5258f2f0be890eff3b7f7679e2cc1b7a	2024-03-01 09:58:25.299887+03	t
5	b5c4d708-71de-4384-a91d-73843bb45947	somesuper@gmail.com	089925fe07a4819e2f032fc4e7b6c1e191dc5d59dbf0bd6b7857a87ef1f7e5c6915ce7ca93f74b136417e362752f81ac5258f2f0be890eff3b7f7679e2cc1b7a	2024-03-12 13:48:50.626919+03	f
10	b9e287e2-d8f4-4083-9efc-d3a0636b3bed	testaddtwo@mail.com	089925fe07a4819e2f032fc4e7b6c1e191dc5d59dbf0bd6b7857a87ef1f7e5c6915ce7ca93f74b136417e362752f81ac5258f2f0be890eff3b7f7679e2cc1b7a	2024-04-19 08:51:53.834211+03	f
\.


--
-- Name: apps_id_seq; Type: SEQUENCE SET; Schema: public; Owner: blueuser
--

SELECT pg_catalog.setval('public.apps_id_seq', 4, false);


--
-- Name: blob_pictures_id_seq; Type: SEQUENCE SET; Schema: public; Owner: blueuser
--

SELECT pg_catalog.setval('public.blob_pictures_id_seq', 1, false);


--
-- Name: blob_videos_id_seq; Type: SEQUENCE SET; Schema: public; Owner: blueuser
--

SELECT pg_catalog.setval('public.blob_videos_id_seq', 1, false);


--
-- Name: end_points_id_seq; Type: SEQUENCE SET; Schema: public; Owner: blueuser
--

SELECT pg_catalog.setval('public.end_points_id_seq', 31, true);


--
-- Name: features_id_seq; Type: SEQUENCE SET; Schema: public; Owner: blueuser
--

SELECT pg_catalog.setval('public.features_id_seq', 1, false);


--
-- Name: jwt_salts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: blueuser
--

SELECT pg_catalog.setval('public.jwt_salts_id_seq', 1, true);


--
-- Name: pages_id_seq; Type: SEQUENCE SET; Schema: public; Owner: blueuser
--

SELECT pg_catalog.setval('public.pages_id_seq', 1, false);


--
-- Name: roles_id_seq; Type: SEQUENCE SET; Schema: public; Owner: blueuser
--

SELECT pg_catalog.setval('public.roles_id_seq', 7, false);


--
-- Name: site_data_id_seq; Type: SEQUENCE SET; Schema: public; Owner: blueuser
--

SELECT pg_catalog.setval('public.site_data_id_seq', 1, false);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: blueuser
--

SELECT pg_catalog.setval('public.users_id_seq', 11, true);


--
-- Name: apps apps_pkey; Type: CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.apps
    ADD CONSTRAINT apps_pkey PRIMARY KEY (id);


--
-- Name: blob_pictures blob_pictures_pkey; Type: CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.blob_pictures
    ADD CONSTRAINT blob_pictures_pkey PRIMARY KEY (id);


--
-- Name: blob_videos blob_videos_pkey; Type: CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.blob_videos
    ADD CONSTRAINT blob_videos_pkey PRIMARY KEY (id);


--
-- Name: end_points end_points_pkey; Type: CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.end_points
    ADD CONSTRAINT end_points_pkey PRIMARY KEY (id);


--
-- Name: features features_pkey; Type: CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.features
    ADD CONSTRAINT features_pkey PRIMARY KEY (id);


--
-- Name: jwt_salts jwt_salts_pkey; Type: CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.jwt_salts
    ADD CONSTRAINT jwt_salts_pkey PRIMARY KEY (id);


--
-- Name: page_roles page_roles_pkey; Type: CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.page_roles
    ADD CONSTRAINT page_roles_pkey PRIMARY KEY (page_id, role_id);


--
-- Name: pages pages_pkey; Type: CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.pages
    ADD CONSTRAINT pages_pkey PRIMARY KEY (id);


--
-- Name: roles roles_pkey; Type: CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_pkey PRIMARY KEY (id);


--
-- Name: site_data site_data_pkey; Type: CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.site_data
    ADD CONSTRAINT site_data_pkey PRIMARY KEY (id);


--
-- Name: apps uni_apps_name; Type: CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.apps
    ADD CONSTRAINT uni_apps_name UNIQUE (name);


--
-- Name: end_points uni_end_points_name; Type: CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.end_points
    ADD CONSTRAINT uni_end_points_name UNIQUE (name);


--
-- Name: features uni_features_name; Type: CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.features
    ADD CONSTRAINT uni_features_name UNIQUE (name);


--
-- Name: pages uni_pages_name; Type: CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.pages
    ADD CONSTRAINT uni_pages_name UNIQUE (name);


--
-- Name: roles uni_roles_description; Type: CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT uni_roles_description UNIQUE (description);


--
-- Name: roles uni_roles_name; Type: CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT uni_roles_name UNIQUE (name);


--
-- Name: users uni_users_email; Type: CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT uni_users_email UNIQUE (email);


--
-- Name: user_roles user_roles_pkey; Type: CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.user_roles
    ADD CONSTRAINT user_roles_pkey PRIMARY KEY (role_id, user_id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_blob_videos_name; Type: INDEX; Schema: public; Owner: blueuser
--

CREATE INDEX idx_blob_videos_name ON public.blob_videos USING btree (name);


--
-- Name: roles fk_apps_roles; Type: FK CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT fk_apps_roles FOREIGN KEY (app_id) REFERENCES public.apps(id);


--
-- Name: end_points fk_features_endpoints; Type: FK CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.end_points
    ADD CONSTRAINT fk_features_endpoints FOREIGN KEY (feature_id) REFERENCES public.features(id);


--
-- Name: page_roles fk_page_roles_page; Type: FK CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.page_roles
    ADD CONSTRAINT fk_page_roles_page FOREIGN KEY (page_id) REFERENCES public.pages(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: page_roles fk_page_roles_role; Type: FK CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.page_roles
    ADD CONSTRAINT fk_page_roles_role FOREIGN KEY (role_id) REFERENCES public.roles(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: features fk_roles_features; Type: FK CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.features
    ADD CONSTRAINT fk_roles_features FOREIGN KEY (role_id) REFERENCES public.roles(id) ON UPDATE CASCADE;


--
-- Name: user_roles fk_user_roles_role; Type: FK CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.user_roles
    ADD CONSTRAINT fk_user_roles_role FOREIGN KEY (role_id) REFERENCES public.roles(id) ON UPDATE CASCADE;


--
-- Name: user_roles fk_user_roles_user; Type: FK CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.user_roles
    ADD CONSTRAINT fk_user_roles_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE;


--
-- PostgreSQL database dump complete
--

