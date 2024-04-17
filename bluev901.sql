--
-- PostgreSQL database dump
--

-- Dumped from database version 15.6 (Ubuntu 15.6-0ubuntu0.23.10.1)
-- Dumped by pg_dump version 16.2 (Ubuntu 16.2-1.pgdg23.10+1)

-- Started on 2024-03-29 13:09:56 EAT

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

--
-- TOC entry 3586 (class 1262 OID 16388)
-- Name: bluev5; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE bluev5 WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.UTF-8';


ALTER DATABASE bluev5 OWNER TO postgres;

\connect bluev5

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
-- TOC entry 219 (class 1259 OID 16413)
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
-- TOC entry 218 (class 1259 OID 16412)
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
-- TOC entry 3588 (class 0 OID 0)
-- Dependencies: 218
-- Name: apps_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: blueuser
--

ALTER SEQUENCE public.apps_id_seq OWNED BY public.apps.id;


--
-- TOC entry 234 (class 1259 OID 16534)
-- Name: blob_pictures; Type: TABLE; Schema: public; Owner: blueuser
--

CREATE TABLE public.blob_pictures (
    id bigint NOT NULL,
    blob_picture bytea
);


ALTER TABLE public.blob_pictures OWNER TO blueuser;

--
-- TOC entry 233 (class 1259 OID 16533)
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
-- TOC entry 3589 (class 0 OID 0)
-- Dependencies: 233
-- Name: blob_pictures_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: blueuser
--

ALTER SEQUENCE public.blob_pictures_id_seq OWNED BY public.blob_pictures.id;


--
-- TOC entry 236 (class 1259 OID 16543)
-- Name: blob_videos; Type: TABLE; Schema: public; Owner: blueuser
--

CREATE TABLE public.blob_videos (
    id bigint NOT NULL,
    name text,
    blob_video bytea
);


ALTER TABLE public.blob_videos OWNER TO blueuser;

--
-- TOC entry 235 (class 1259 OID 16542)
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
-- TOC entry 3590 (class 0 OID 0)
-- Dependencies: 235
-- Name: blob_videos_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: blueuser
--

ALTER SEQUENCE public.blob_videos_id_seq OWNED BY public.blob_videos.id;


--
-- TOC entry 229 (class 1259 OID 16503)
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
-- TOC entry 228 (class 1259 OID 16502)
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
-- TOC entry 3591 (class 0 OID 0)
-- Dependencies: 228
-- Name: end_points_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: blueuser
--

ALTER SEQUENCE public.end_points_id_seq OWNED BY public.end_points.id;


--
-- TOC entry 227 (class 1259 OID 16486)
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
-- TOC entry 226 (class 1259 OID 16485)
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
-- TOC entry 3592 (class 0 OID 0)
-- Dependencies: 226
-- Name: features_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: blueuser
--

ALTER SEQUENCE public.features_id_seq OWNED BY public.features.id;


--
-- TOC entry 215 (class 1259 OID 16391)
-- Name: jwt_salts; Type: TABLE; Schema: public; Owner: blueuser
--

CREATE TABLE public.jwt_salts (
    id bigint NOT NULL,
    salt_a text,
    salt_b text
);


ALTER TABLE public.jwt_salts OWNER TO blueuser;

--
-- TOC entry 214 (class 1259 OID 16390)
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
-- TOC entry 3593 (class 0 OID 0)
-- Dependencies: 214
-- Name: jwt_salts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: blueuser
--

ALTER SEQUENCE public.jwt_salts_id_seq OWNED BY public.jwt_salts.id;


--
-- TOC entry 225 (class 1259 OID 16470)
-- Name: page_roles; Type: TABLE; Schema: public; Owner: blueuser
--

CREATE TABLE public.page_roles (
    page_id bigint NOT NULL,
    role_id bigint NOT NULL
);


ALTER TABLE public.page_roles OWNER TO blueuser;

--
-- TOC entry 224 (class 1259 OID 16459)
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
-- TOC entry 223 (class 1259 OID 16458)
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
-- TOC entry 3594 (class 0 OID 0)
-- Dependencies: 223
-- Name: pages_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: blueuser
--

ALTER SEQUENCE public.pages_id_seq OWNED BY public.pages.id;


--
-- TOC entry 221 (class 1259 OID 16425)
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
-- TOC entry 220 (class 1259 OID 16424)
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
-- TOC entry 3595 (class 0 OID 0)
-- Dependencies: 220
-- Name: roles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: blueuser
--

ALTER SEQUENCE public.roles_id_seq OWNED BY public.roles.id;


--
-- TOC entry 230 (class 1259 OID 16518)
-- Name: session_data; Type: TABLE; Schema: public; Owner: blueuser
--

CREATE TABLE public.session_data (
    token text,
    time_stamp timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.session_data OWNER TO blueuser;

--
-- TOC entry 232 (class 1259 OID 16525)
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
-- TOC entry 231 (class 1259 OID 16524)
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
-- TOC entry 3596 (class 0 OID 0)
-- Dependencies: 231
-- Name: site_data_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: blueuser
--

ALTER SEQUENCE public.site_data_id_seq OWNED BY public.site_data.id;


--
-- TOC entry 222 (class 1259 OID 16443)
-- Name: user_roles; Type: TABLE; Schema: public; Owner: blueuser
--

CREATE TABLE public.user_roles (
    role_id bigint NOT NULL,
    user_id bigint NOT NULL
);


ALTER TABLE public.user_roles OWNER TO blueuser;

--
-- TOC entry 217 (class 1259 OID 16400)
-- Name: users; Type: TABLE; Schema: public; Owner: blueuser
--

CREATE TABLE public.users (
    id SERIAL PRIMARY KEY,
    uuid UUID,
    email TEXT,
    password TEXT,
    date_registered TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    disabled BOOLEAN DEFAULT FALSE
);


ALTER TABLE public.users OWNER TO blueuser;

--
-- TOC entry 216 (class 1259 OID 16399)
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
-- TOC entry 3597 (class 0 OID 0)
-- Dependencies: 216
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: blueuser
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 3357 (class 2604 OID 16416)
-- Name: apps id; Type: DEFAULT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.apps ALTER COLUMN id SET DEFAULT nextval('public.apps_id_seq'::regclass);


--
-- TOC entry 3368 (class 2604 OID 16537)
-- Name: blob_pictures id; Type: DEFAULT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.blob_pictures ALTER COLUMN id SET DEFAULT nextval('public.blob_pictures_id_seq'::regclass);


--
-- TOC entry 3369 (class 2604 OID 16546)
-- Name: blob_videos id; Type: DEFAULT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.blob_videos ALTER COLUMN id SET DEFAULT nextval('public.blob_videos_id_seq'::regclass);


--
-- TOC entry 3365 (class 2604 OID 16506)
-- Name: end_points id; Type: DEFAULT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.end_points ALTER COLUMN id SET DEFAULT nextval('public.end_points_id_seq'::regclass);


--
-- TOC entry 3363 (class 2604 OID 16489)
-- Name: features id; Type: DEFAULT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.features ALTER COLUMN id SET DEFAULT nextval('public.features_id_seq'::regclass);


--
-- TOC entry 3353 (class 2604 OID 16394)
-- Name: jwt_salts id; Type: DEFAULT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.jwt_salts ALTER COLUMN id SET DEFAULT nextval('public.jwt_salts_id_seq'::regclass);


--
-- TOC entry 3361 (class 2604 OID 16462)
-- Name: pages id; Type: DEFAULT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.pages ALTER COLUMN id SET DEFAULT nextval('public.pages_id_seq'::regclass);


--
-- TOC entry 3359 (class 2604 OID 16428)
-- Name: roles id; Type: DEFAULT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.roles ALTER COLUMN id SET DEFAULT nextval('public.roles_id_seq'::regclass);


--
-- TOC entry 3367 (class 2604 OID 16528)
-- Name: site_data id; Type: DEFAULT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.site_data ALTER COLUMN id SET DEFAULT nextval('public.site_data_id_seq'::regclass);


--
-- TOC entry 3354 (class 2604 OID 16403)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 3563 (class 0 OID 16413)
-- Dependencies: 219
-- Data for Name: apps; Type: TABLE DATA; Schema: public; Owner: blueuser
--

INSERT INTO public.apps (id, name, uuid, active, description) VALUES
	(1, 'BlueAdmin', '48015a9b-5a86-4a15-944b-94108aa78b4b', true, 'SSO Role Based User Administration solution'),
	(2, 'BlueCom', '21028fa1-8e04-464e-8be7-95ea9c82994b', true, 'Commercial Order Management'),
	(3, 'BlueHRM', '9359b1ba-98b5-427c-96d8-023fc33cd1b0', true, 'Human Resource Management') ON CONFLICT DO NOTHING;


--
-- TOC entry 3578 (class 0 OID 16534)
-- Dependencies: 234
-- Data for Name: blob_pictures; Type: TABLE DATA; Schema: public; Owner: blueuser
--



--
-- TOC entry 3580 (class 0 OID 16543)
-- Dependencies: 236
-- Data for Name: blob_videos; Type: TABLE DATA; Schema: public; Owner: blueuser
--



--
-- TOC entry 3573 (class 0 OID 16503)
-- Dependencies: 229
-- Data for Name: end_points; Type: TABLE DATA; Schema: public; Owner: blueuser
--

INSERT INTO public.end_points (id, name, route_paths, method, description, feature_id) VALUES
	(1, 'swagger_routes_get', '/docs/*', 'GET', 'swagger_routes-GET', NULL),
	(2, 'custom_metrics_route_get', '/lmetrics', 'GET', 'custom_metrics_route-GET', 13),
	(3, 'check_login_get', '/api/v1/checklogin', 'GET', 'check_login-GET', 13),
	(4, 'roles_get', '/api/v1/roles', 'GET', 'roles-GET', 1),
	(5, 'roles_single_get', '/api/v1/roles/:id', 'GET', 'roles_single-GET', 1),
	(6, 'drop_roles_get', '/api/v1/droproles', 'GET', 'drop_roles-GET', 14),
	(7, 'roles_endpoints_get', '/api/v1/role_endpoints', 'GET', 'roles_endpoints-GET', 1),
	(8, 'features_get', '/api/v1/features', 'GET', 'features-GET', 11),
	(9, 'features_single_get', '/api/v1/features/:id', 'GET', 'features_single-GET', 11),
	(10, 'drop_features_get', '/api/v1/featuredrop', 'GET', 'drop_features-GET', 14),
	(11, 'apps_get', '/api/v1/apps', 'GET', 'apps-GET', 7),
	(12, 'apps_single_get', '/api/v1/apps/:id', 'GET', 'apps_single-GET', 7),
	(13, 'drop_sppd_get', '/api/v1/appsdrop', 'GET', 'drop_sppd-GET', 14),
	(14, 'apps_features_get', '/api/v1/appsmatrix/:id', 'GET', 'apps_features-GET', 7),
	(15, 'users_get', '/api/v1/users', 'GET', 'users-GET', 3),
	(16, 'user_single_get', '/api/v1/users/:id', 'GET', 'user_single-GET', 3),
	(17, 'get_user_roles_get', '/api/v1/userrole/:user_id', 'GET', 'get_user_roles-GET', 3),
	(18, 'pages_get', '/api/v1/pages', 'GET', 'pages-GET', 5),
	(19, 'page_single_get', '/api/v1/pages/:id', 'GET', 'page_single-GET', 5),
	(20, 'get_page_roles_get', '/api/v1/pagesroles/:page_id', 'GET', 'get_page_roles-GET', 5),
	(21, 'end_point_get', '/api/v1/endpoints', 'GET', 'end_point-GET', 9),
	(22, 'end_point_single_get', '/api/v1/endpoints/:id', 'GET', 'end_point_single-GET', 9),
	(23, 'drop_endpoints_get', '/api/v1/endpointdrop', 'GET', 'drop_endpoints-GET', 14),
	(25, 'login_route_post', '/api/v1/login', 'POST', 'login_route-POST', 13),
	(26, 'roles_post', '/api/v1/roles', 'POST', 'roles-POST', 2),
	(27, 'features_post', '/api/v1/features', 'POST', 'features-POST', 12),
	(28, 'apps_post', '/api/v1/apps', 'POST', 'apps-POST', 8),
	(29, 'users_post', '/api/v1/users', 'POST', 'users-POST', 4),
	(30, 'user_role_post', '/api/v1/userrole/:user_id/:role_id', 'POST', 'user_role-POST', 4),
	(31, 'pages_post', '/api/v1/pages', 'POST', 'pages-POST', 6),
	(32, 'page_roles_post', '/api/v1/pagerole/:page_id/:role_id', 'POST', 'page_roles-POST', 6),
	(33, 'end_point_post', '/api/v1/endpoints', 'POST', 'end_point-POST', 10),
	(34, 'send_email_post', '/api/v1/email', 'POST', 'send_email-POST', NULL),
	(35, 'blob_picture_post', '/api/v1/blobpic', 'POST', 'blob_picture-POST', 13),
	(36, 'blob_video_post', '/api/v1/blobvideo', 'POST', 'blob_video-POST', 13),
	(37, 'activate_deactivate_role_put', '/api/v1/roles/:role_id', 'PUT', 'activate_deactivate_role-PUT', 2),
	(38, 'activate_deactivate_features_put', '/api/v1/features/:feature_id', 'PUT', 'activate_deactivate_features-PUT', 12),
	(39, 'activate_deactivate_user_put', '/api/v1/users/:user_id', 'PUT', 'activate_deactivate_user-PUT', 4),
	(40, 'roles_single_delete', '/api/v1/roles/:id', 'DELETE', 'roles_single-DELETE', 2),
	(41, 'features_single_delete', '/api/v1/features/:id', 'DELETE', 'features_single-DELETE', 12),
	(42, 'feature_role_delete', '/api/v1/featuresrole/:feature_id', 'DELETE', 'feature_role-DELETE', 12),
	(43, 'apps_single_delete', '/api/v1/apps/:id', 'DELETE', 'apps_single-DELETE', 8),
	(44, 'user_single_delete', '/api/v1/users/:id', 'DELETE', 'user_single-DELETE', 4),
	(45, 'user_role_delete', '/api/v1/userrole/:user_id/:role_id', 'DELETE', 'user_role-DELETE', 4),
	(46, 'page_single_delete', '/api/v1/pages/:id', 'DELETE', 'page_single-DELETE', 6),
	(47, 'page_roles_delete', '/api/v1/pagerole/:page_id/:role_id', 'DELETE', 'page_roles-DELETE', 6),
	(48, 'end_point_single_delete', '/api/v1/endpoints/:id', 'DELETE', 'end_point_single-DELETE', 10),
	(49, 'feature_endpoint_delete', '/api/v1/feature_endpoint/:endpoint_id', 'DELETE', 'feature_endpoint-DELETE', 12),
	(50, 'roles_single_patch', '/api/v1/roles/:id', 'PATCH', 'roles_single-PATCH', 2),
	(51, 'roles_app_patch', '/api/v1/approle/:role_id', 'PATCH', 'roles_app-PATCH', 2),
	(52, 'features_single_patch', '/api/v1/features/:id', 'PATCH', 'features_single-PATCH', 12),
	(53, 'feature_role_patch', '/api/v1/featuresrole/:feature_id', 'PATCH', 'feature_role-PATCH', 12),
	(54, 'apps_single_patch', '/api/v1/apps/:id', 'PATCH', 'apps_single-PATCH', 8),
	(55, 'user_single_patch', '/api/v1/users/:id', 'PATCH', 'user_single-PATCH', 4),
	(56, 'page_single_patch', '/api/v1/pages/:id', 'PATCH', 'page_single-PATCH', 6),
	(57, 'end_point_single_patch', '/api/v1/endpoints/:id', 'PATCH', 'end_point_single-PATCH', 10),
	(58, 'feature_endpoint_patch', '/api/v1/feature_endpoint/:endpoint_id', 'PATCH', 'feature_endpoint-PATCH', 12),
	(59, 'change_reset_password_put', '/api/v1/users/:email_id', 'PUT', 'change_reset_password-PUT', 4),
	(24, 'dashboard_get', '/api/v1/dashboard', 'GET', 'dashboard-GET', 13) ON CONFLICT DO NOTHING;


--
-- TOC entry 3571 (class 0 OID 16486)
-- Dependencies: 227
-- Data for Name: features; Type: TABLE DATA; Schema: public; Owner: blueuser
--

INSERT INTO public.features (id, name, description, active, role_id) VALUES
	(1, 'role_read', 'View List of Roles', true, 2),
	(2, 'role_write', 'Privilege to Update ,Create, View and Delete  ', true, 3),
	(3, 'user_read', 'View List of Users', true, 2),
	(4, 'user_write', 'Privilege to Update ,Create, View and Delete  ', true, 3),
	(5, 'page_read', 'View List of Pages', true, 2),
	(6, 'page_write', 'Privilege to Update ,Create, View and Delete  ', true, 3),
	(7, 'app_read', 'View List of Apps', true, 2),
	(8, 'app_write', 'Privilege to Update ,Create, View and Delete ', true, 4),
	(9, 'endpoint_read', 'View List of Endpoints', true, 2),
	(10, 'endpoint_write', 'Privilege to Update ,Create, View and Delete  ', true, 3),
	(11, 'feature_read', 'View List of Features', true, 2),
	(12, 'features_write', 'Privilege to Update ,Create, View and Delete  ', true, 3),
	(13, 'login', 'Endpoints that can be accessed with out Logging In', true, 5),
	(14, 'drop_down', 'Endpoints that fetch Drowns ', true, 6) ON CONFLICT DO NOTHING;


--
-- TOC entry 3559 (class 0 OID 16391)
-- Dependencies: 215
-- Data for Name: jwt_salts; Type: TABLE DATA; Schema: public; Owner: blueuser
--

INSERT INTO public.jwt_salts (id, salt_a, salt_b) VALUES
	(1, 'jwCNUyNjLw3fm3BtGZNnSoNjn', 'L8PaWn1QcU3OQfw3Sh3dIFr4j') ON CONFLICT DO NOTHING;


--
-- TOC entry 3569 (class 0 OID 16470)
-- Dependencies: 225
-- Data for Name: page_roles; Type: TABLE DATA; Schema: public; Owner: blueuser
--

INSERT INTO public.page_roles (page_id, role_id) VALUES
	(1, 5),
	(5, 3),
	(2, 5),
	(3, 3),
	(3, 1),
	(2, 1),
	(1, 1),
	(4, 2),
	(4, 3),
	(4, 1),
	(5, 2),
	(5, 1),
	(6, 1),
	(6, 2),
	(6, 3),
	(7, 1),
	(7, 3),
	(8, 1),
	(8, 3),
	(7, 2),
	(3, 2),
	(8, 2),
	(2, 2),
	(9, 1) ON CONFLICT DO NOTHING;


--
-- TOC entry 3568 (class 0 OID 16459)
-- Dependencies: 224
-- Data for Name: pages; Type: TABLE DATA; Schema: public; Owner: blueuser
--

INSERT INTO public.pages (id, name, active, description) VALUES
	(1, 'Login', true, 'Login'),
	(2, 'Home', true, 'Home'),
	(3, 'Role', true, 'View Roles'),
	(4, 'User', true, 'View Users'),
	(5, 'Page', true, 'View Pages'),
	(6, 'Feature', true, 'View Features'),
	(7, 'Endpoint', true, 'View Endpoints'),
	(8, 'App', true, 'View Apps'),
	(9, 'Sign Up', true, 'Sign Up Page For users Who have not registered') ON CONFLICT DO NOTHING;


--
-- TOC entry 3565 (class 0 OID 16425)
-- Dependencies: 221
-- Data for Name: roles; Type: TABLE DATA; Schema: public; Owner: blueuser
--

INSERT INTO public.roles (id, name, description, active, app_id) VALUES
	(1, 'superuser', 'Have Access to All resources in the Apps', true, 1),
	(2, 'standard', 'Have Access to Limited Features', true, 1),
	(3, 'administrator', 'Have Access System Admin Privileges', true, 1),
	(4, 'app_role', 'Have App CURD abilities', true, 1),
	(5, 'Anonymous', 'For Pages that do not need user sign in', true, 1),
	(6, 'Drop Down ', 'To access Drop Down menu fetching  endpoints', true, 1) ON CONFLICT DO NOTHING;


--
-- TOC entry 3574 (class 0 OID 16518)
-- Dependencies: 230
-- Data for Name: session_data; Type: TABLE DATA; Schema: public; Owner: blueuser
--



--
-- TOC entry 3576 (class 0 OID 16525)
-- Dependencies: 232
-- Data for Name: site_data; Type: TABLE DATA; Schema: public; Owner: blueuser
--



--
-- TOC entry 3566 (class 0 OID 16443)
-- Dependencies: 222
-- Data for Name: user_roles; Type: TABLE DATA; Schema: public; Owner: blueuser
--

INSERT INTO public.user_roles (role_id, user_id) VALUES
	(1, 1),
	(1, 4),
	(1, 5),
	(2, 2),
	(2, 6),
	(2, 7),
	(3, 3) ON CONFLICT DO NOTHING;


--
-- TOC entry 3561 (class 0 OID 16400)
-- Dependencies: 217
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: blueuser
--

INSERT INTO public.users (id, uuid, email, password, date_registered, disabled) VALUES
	(4, '8a200cd4-9067-4508-b93c-4b242ef03740', 'beimnet.degefu@gmail.com', '089925fe07a4819e2f032fc4e7b6c1e191dc5d59dbf0bd6b7857a87ef1f7e5c6915ce7ca93f74b136417e362752f81ac5258f2f0be890eff3b7f7679e2cc1b7a', '2024-03-08 19:49:53.754679+03', false),
	(6, '79731de0-de10-4f62-bcc5-7637d132110a', 'mickyasne123@gmail.com', '3d5267ceaa0759b8756837e4e817517fae8fa2d168267053e78428ac0941d5542c7bee61536be21f7d117fed80bc9aa27e2d450e261ab361ece2a701a1e2da72', '2024-03-12 13:52:01.052905+03', false),
	(7, '586c2bc7-9a66-49df-a81e-93b2efaec8c9', 'mickyasne12@gmail.com', 'adfebcc24b8ce4be96b83381f08275bbd6fe355dee6a1a1407247cd15bbb8ec5952b5680af51c532d8785f950e4bb7ddeab72efc737c4656bcd27b7f73b72bf1', '2024-03-12 15:03:48.01416+03', false),
	(1, '38ca7360-0138-4b0f-8985-b307ad188e92', 'superuser@mail.com', '089925fe07a4819e2f032fc4e7b6c1e191dc5d59dbf0bd6b7857a87ef1f7e5c6915ce7ca93f74b136417e362752f81ac5258f2f0be890eff3b7f7679e2cc1b7a', '2024-01-16 11:27:55.628203+03', false),
	(2, '1323b2d9-5755-4e4c-9af5-93c17f59e6fd', 'standarduser@mail.com', '089925fe07a4819e2f032fc4e7b6c1e191dc5d59dbf0bd6b7857a87ef1f7e5c6915ce7ca93f74b136417e362752f81ac5258f2f0be890eff3b7f7679e2cc1b7a', '2024-01-17 15:33:56.443849+03', false),
	(3, '12bc4954-487b-4263-8bdc-cacbf720f623', 'adminuser@mail.com', '089925fe07a4819e2f032fc4e7b6c1e191dc5d59dbf0bd6b7857a87ef1f7e5c6915ce7ca93f74b136417e362752f81ac5258f2f0be890eff3b7f7679e2cc1b7a', '2024-03-01 09:58:25.299887+03', true),
	(5, 'b5c4d708-71de-4384-a91d-73843bb45947', 'somesuper@gmail.com', '089925fe07a4819e2f032fc4e7b6c1e191dc5d59dbf0bd6b7857a87ef1f7e5c6915ce7ca93f74b136417e362752f81ac5258f2f0be890eff3b7f7679e2cc1b7a', '2024-03-12 13:48:50.626919+03', false) ON CONFLICT DO NOTHING;


--
-- TOC entry 3598 (class 0 OID 0)
-- Dependencies: 218
-- Name: apps_id_seq; Type: SEQUENCE SET; Schema: public; Owner: blueuser
--

SELECT pg_catalog.setval('public.apps_id_seq', 1, false);


--
-- TOC entry 3599 (class 0 OID 0)
-- Dependencies: 233
-- Name: blob_pictures_id_seq; Type: SEQUENCE SET; Schema: public; Owner: blueuser
--

SELECT pg_catalog.setval('public.blob_pictures_id_seq', 1, false);


--
-- TOC entry 3600 (class 0 OID 0)
-- Dependencies: 235
-- Name: blob_videos_id_seq; Type: SEQUENCE SET; Schema: public; Owner: blueuser
--

SELECT pg_catalog.setval('public.blob_videos_id_seq', 1, false);


--
-- TOC entry 3601 (class 0 OID 0)
-- Dependencies: 228
-- Name: end_points_id_seq; Type: SEQUENCE SET; Schema: public; Owner: blueuser
--

SELECT pg_catalog.setval('public.end_points_id_seq', 2043, true);


--
-- TOC entry 3602 (class 0 OID 0)
-- Dependencies: 226
-- Name: features_id_seq; Type: SEQUENCE SET; Schema: public; Owner: blueuser
--

SELECT pg_catalog.setval('public.features_id_seq', 1, false);


--
-- TOC entry 3603 (class 0 OID 0)
-- Dependencies: 214
-- Name: jwt_salts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: blueuser
--

SELECT pg_catalog.setval('public.jwt_salts_id_seq', 1, true);


--
-- TOC entry 3604 (class 0 OID 0)
-- Dependencies: 223
-- Name: pages_id_seq; Type: SEQUENCE SET; Schema: public; Owner: blueuser
--

SELECT pg_catalog.setval('public.pages_id_seq', 1, false);


--
-- TOC entry 3605 (class 0 OID 0)
-- Dependencies: 220
-- Name: roles_id_seq; Type: SEQUENCE SET; Schema: public; Owner: blueuser
--

SELECT pg_catalog.setval('public.roles_id_seq', 1, false);


--
-- TOC entry 3606 (class 0 OID 0)
-- Dependencies: 231
-- Name: site_data_id_seq; Type: SEQUENCE SET; Schema: public; Owner: blueuser
--

SELECT pg_catalog.setval('public.site_data_id_seq', 1, false);


--
-- TOC entry 3607 (class 0 OID 0)
-- Dependencies: 216
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: blueuser
--

SELECT pg_catalog.setval('public.users_id_seq', 1, false);


--
-- TOC entry 3377 (class 2606 OID 16421)
-- Name: apps apps_pkey; Type: CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.apps
    ADD CONSTRAINT apps_pkey PRIMARY KEY (id);


--
-- TOC entry 3405 (class 2606 OID 16541)
-- Name: blob_pictures blob_pictures_pkey; Type: CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.blob_pictures
    ADD CONSTRAINT blob_pictures_pkey PRIMARY KEY (id);


--
-- TOC entry 3407 (class 2606 OID 16550)
-- Name: blob_videos blob_videos_pkey; Type: CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.blob_videos
    ADD CONSTRAINT blob_videos_pkey PRIMARY KEY (id);


--
-- TOC entry 3399 (class 2606 OID 16510)
-- Name: end_points end_points_pkey; Type: CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.end_points
    ADD CONSTRAINT end_points_pkey PRIMARY KEY (id);


--
-- TOC entry 3395 (class 2606 OID 16494)
-- Name: features features_pkey; Type: CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.features
    ADD CONSTRAINT features_pkey PRIMARY KEY (id);


--
-- TOC entry 3371 (class 2606 OID 16398)
-- Name: jwt_salts jwt_salts_pkey; Type: CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.jwt_salts
    ADD CONSTRAINT jwt_salts_pkey PRIMARY KEY (id);


--
-- TOC entry 3393 (class 2606 OID 16474)
-- Name: page_roles page_roles_pkey; Type: CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.page_roles
    ADD CONSTRAINT page_roles_pkey PRIMARY KEY (page_id, role_id);


--
-- TOC entry 3389 (class 2606 OID 16467)
-- Name: pages pages_pkey; Type: CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.pages
    ADD CONSTRAINT pages_pkey PRIMARY KEY (id);


--
-- TOC entry 3381 (class 2606 OID 16433)
-- Name: roles roles_pkey; Type: CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_pkey PRIMARY KEY (id);


--
-- TOC entry 3403 (class 2606 OID 16532)
-- Name: site_data site_data_pkey; Type: CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.site_data
    ADD CONSTRAINT site_data_pkey PRIMARY KEY (id);


--
-- TOC entry 3379 (class 2606 OID 16423)
-- Name: apps uni_apps_name; Type: CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.apps
    ADD CONSTRAINT uni_apps_name UNIQUE (name);


--
-- TOC entry 3401 (class 2606 OID 16512)
-- Name: end_points uni_end_points_name; Type: CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.end_points
    ADD CONSTRAINT uni_end_points_name UNIQUE (name);


--
-- TOC entry 3397 (class 2606 OID 16496)
-- Name: features uni_features_name; Type: CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.features
    ADD CONSTRAINT uni_features_name UNIQUE (name);


--
-- TOC entry 3391 (class 2606 OID 16469)
-- Name: pages uni_pages_name; Type: CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.pages
    ADD CONSTRAINT uni_pages_name UNIQUE (name);


--
-- TOC entry 3383 (class 2606 OID 16437)
-- Name: roles uni_roles_description; Type: CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT uni_roles_description UNIQUE (description);


--
-- TOC entry 3385 (class 2606 OID 16435)
-- Name: roles uni_roles_name; Type: CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT uni_roles_name UNIQUE (name);


--
-- TOC entry 3373 (class 2606 OID 16411)
-- Name: users uni_users_email; Type: CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT uni_users_email UNIQUE (email);


--
-- TOC entry 3387 (class 2606 OID 16447)
-- Name: user_roles user_roles_pkey; Type: CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.user_roles
    ADD CONSTRAINT user_roles_pkey PRIMARY KEY (role_id, user_id);


--
-- TOC entry 3375 (class 2606 OID 16409)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 3408 (class 1259 OID 16551)
-- Name: idx_blob_videos_name; Type: INDEX; Schema: public; Owner: blueuser
--

CREATE INDEX idx_blob_videos_name ON public.blob_videos USING btree (name);


--
-- TOC entry 3409 (class 2606 OID 16438)
-- Name: roles fk_apps_roles; Type: FK CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT fk_apps_roles FOREIGN KEY (app_id) REFERENCES public.apps(id);


--
-- TOC entry 3415 (class 2606 OID 16513)
-- Name: end_points fk_features_endpoints; Type: FK CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.end_points
    ADD CONSTRAINT fk_features_endpoints FOREIGN KEY (feature_id) REFERENCES public.features(id);


--
-- TOC entry 3412 (class 2606 OID 16475)
-- Name: page_roles fk_page_roles_page; Type: FK CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.page_roles
    ADD CONSTRAINT fk_page_roles_page FOREIGN KEY (page_id) REFERENCES public.pages(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- TOC entry 3413 (class 2606 OID 16480)
-- Name: page_roles fk_page_roles_role; Type: FK CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.page_roles
    ADD CONSTRAINT fk_page_roles_role FOREIGN KEY (role_id) REFERENCES public.roles(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- TOC entry 3414 (class 2606 OID 16497)
-- Name: features fk_roles_features; Type: FK CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.features
    ADD CONSTRAINT fk_roles_features FOREIGN KEY (role_id) REFERENCES public.roles(id) ON UPDATE CASCADE;


--
-- TOC entry 3410 (class 2606 OID 16448)
-- Name: user_roles fk_user_roles_role; Type: FK CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.user_roles
    ADD CONSTRAINT fk_user_roles_role FOREIGN KEY (role_id) REFERENCES public.roles(id) ON UPDATE CASCADE;


--
-- TOC entry 3411 (class 2606 OID 16453)
-- Name: user_roles fk_user_roles_user; Type: FK CONSTRAINT; Schema: public; Owner: blueuser
--

ALTER TABLE ONLY public.user_roles
    ADD CONSTRAINT fk_user_roles_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE;


--
-- TOC entry 3587 (class 0 OID 0)
-- Dependencies: 3586
-- Name: DATABASE bluev5; Type: ACL; Schema: -; Owner: postgres
--

GRANT ALL ON DATABASE bluev5 TO blueuser;


-- Completed on 2024-03-29 13:09:57 EAT

--
-- PostgreSQL database dump complete
--

