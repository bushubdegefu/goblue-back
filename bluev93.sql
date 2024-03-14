PGDMP      :                |            bluev52    16.2     16.2 (Ubuntu 16.2-1.pgdg23.10+1) Y    o           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            p           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            q           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            r           1262    40960    bluev52    DATABASE     i   CREATE DATABASE bluev52 WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'C';
    DROP DATABASE bluev52;
                bushubdegefu    false            s           0    0    DATABASE bluev52    ACL     1   GRANT ALL ON DATABASE bluev52 TO neon_superuser;
                   bushubdegefu    false    3442            �            1259    49166    apps    TABLE     �   CREATE TABLE public.apps (
    id bigint NOT NULL,
    name text,
    uuid uuid,
    active boolean DEFAULT true,
    description text
);
    DROP TABLE public.apps;
       public         heap    bushubdegefu    false            �            1259    49165    apps_id_seq    SEQUENCE     t   CREATE SEQUENCE public.apps_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 "   DROP SEQUENCE public.apps_id_seq;
       public          bushubdegefu    false    218            t           0    0    apps_id_seq    SEQUENCE OWNED BY     ;   ALTER SEQUENCE public.apps_id_seq OWNED BY public.apps.id;
          public          bushubdegefu    false    217            �            1259    49272    blob_pictures    TABLE     V   CREATE TABLE public.blob_pictures (
    id bigint NOT NULL,
    blob_picture bytea
);
 !   DROP TABLE public.blob_pictures;
       public         heap    bushubdegefu    false            �            1259    49271    blob_pictures_id_seq    SEQUENCE     }   CREATE SEQUENCE public.blob_pictures_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 +   DROP SEQUENCE public.blob_pictures_id_seq;
       public          bushubdegefu    false    230            u           0    0    blob_pictures_id_seq    SEQUENCE OWNED BY     M   ALTER SEQUENCE public.blob_pictures_id_seq OWNED BY public.blob_pictures.id;
          public          bushubdegefu    false    229            �            1259    49281    blob_videos    TABLE     a   CREATE TABLE public.blob_videos (
    id bigint NOT NULL,
    name text,
    blob_video bytea
);
    DROP TABLE public.blob_videos;
       public         heap    bushubdegefu    false            �            1259    49280    blob_videos_id_seq    SEQUENCE     {   CREATE SEQUENCE public.blob_videos_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 )   DROP SEQUENCE public.blob_videos_id_seq;
       public          bushubdegefu    false    232            v           0    0    blob_videos_id_seq    SEQUENCE OWNED BY     I   ALTER SEQUENCE public.blob_videos_id_seq OWNED BY public.blob_videos.id;
          public          bushubdegefu    false    231            �            1259    49256 
   end_points    TABLE     �   CREATE TABLE public.end_points (
    id bigint NOT NULL,
    name text,
    route_paths text,
    method text,
    description text,
    feature_id bigint
);
    DROP TABLE public.end_points;
       public         heap    bushubdegefu    false            �            1259    49255    end_points_id_seq    SEQUENCE     z   CREATE SEQUENCE public.end_points_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 (   DROP SEQUENCE public.end_points_id_seq;
       public          bushubdegefu    false    228            w           0    0    end_points_id_seq    SEQUENCE OWNED BY     G   ALTER SEQUENCE public.end_points_id_seq OWNED BY public.end_points.id;
          public          bushubdegefu    false    227            �            1259    49239    features    TABLE     �   CREATE TABLE public.features (
    id bigint NOT NULL,
    name text,
    description text,
    active boolean DEFAULT true,
    role_id bigint
);
    DROP TABLE public.features;
       public         heap    bushubdegefu    false            �            1259    49238    features_id_seq    SEQUENCE     x   CREATE SEQUENCE public.features_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 &   DROP SEQUENCE public.features_id_seq;
       public          bushubdegefu    false    226            x           0    0    features_id_seq    SEQUENCE OWNED BY     C   ALTER SEQUENCE public.features_id_seq OWNED BY public.features.id;
          public          bushubdegefu    false    225            �            1259    49291 	   jwt_salts    TABLE     \   CREATE TABLE public.jwt_salts (
    id bigint NOT NULL,
    salt_a text,
    salt_b text
);
    DROP TABLE public.jwt_salts;
       public         heap    bushubdegefu    false            �            1259    49290    jwt_salts_id_seq    SEQUENCE     y   CREATE SEQUENCE public.jwt_salts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 '   DROP SEQUENCE public.jwt_salts_id_seq;
       public          bushubdegefu    false    234            y           0    0    jwt_salts_id_seq    SEQUENCE OWNED BY     E   ALTER SEQUENCE public.jwt_salts_id_seq OWNED BY public.jwt_salts.id;
          public          bushubdegefu    false    233            �            1259    49223 
   page_roles    TABLE     ]   CREATE TABLE public.page_roles (
    page_id bigint NOT NULL,
    role_id bigint NOT NULL
);
    DROP TABLE public.page_roles;
       public         heap    bushubdegefu    false            �            1259    49212    pages    TABLE     |   CREATE TABLE public.pages (
    id bigint NOT NULL,
    name text,
    active boolean DEFAULT true,
    description text
);
    DROP TABLE public.pages;
       public         heap    bushubdegefu    false            �            1259    49211    pages_id_seq    SEQUENCE     u   CREATE SEQUENCE public.pages_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.pages_id_seq;
       public          bushubdegefu    false    223            z           0    0    pages_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.pages_id_seq OWNED BY public.pages.id;
          public          bushubdegefu    false    222            �            1259    49178    roles    TABLE     �   CREATE TABLE public.roles (
    id bigint NOT NULL,
    name text NOT NULL,
    description text NOT NULL,
    active boolean DEFAULT true,
    app_id bigint
);
    DROP TABLE public.roles;
       public         heap    bushubdegefu    false            �            1259    49177    roles_id_seq    SEQUENCE     u   CREATE SEQUENCE public.roles_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.roles_id_seq;
       public          bushubdegefu    false    220            {           0    0    roles_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.roles_id_seq OWNED BY public.roles.id;
          public          bushubdegefu    false    219            �            1259    49196 
   user_roles    TABLE     ]   CREATE TABLE public.user_roles (
    role_id bigint NOT NULL,
    user_id bigint NOT NULL
);
    DROP TABLE public.user_roles;
       public         heap    bushubdegefu    false            �            1259    49153    users    TABLE     �   CREATE TABLE public.users (
    id bigint NOT NULL,
    uuid uuid,
    email text,
    password text,
    date_registered timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    disabled boolean DEFAULT false
);
    DROP TABLE public.users;
       public         heap    bushubdegefu    false            �            1259    49152    users_id_seq    SEQUENCE     u   CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.users_id_seq;
       public          bushubdegefu    false    216            |           0    0    users_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;
          public          bushubdegefu    false    215            �           2604    49169    apps id    DEFAULT     b   ALTER TABLE ONLY public.apps ALTER COLUMN id SET DEFAULT nextval('public.apps_id_seq'::regclass);
 6   ALTER TABLE public.apps ALTER COLUMN id DROP DEFAULT;
       public          bushubdegefu    false    217    218    218            �           2604    49275    blob_pictures id    DEFAULT     t   ALTER TABLE ONLY public.blob_pictures ALTER COLUMN id SET DEFAULT nextval('public.blob_pictures_id_seq'::regclass);
 ?   ALTER TABLE public.blob_pictures ALTER COLUMN id DROP DEFAULT;
       public          bushubdegefu    false    230    229    230            �           2604    49284    blob_videos id    DEFAULT     p   ALTER TABLE ONLY public.blob_videos ALTER COLUMN id SET DEFAULT nextval('public.blob_videos_id_seq'::regclass);
 =   ALTER TABLE public.blob_videos ALTER COLUMN id DROP DEFAULT;
       public          bushubdegefu    false    232    231    232            �           2604    49259    end_points id    DEFAULT     n   ALTER TABLE ONLY public.end_points ALTER COLUMN id SET DEFAULT nextval('public.end_points_id_seq'::regclass);
 <   ALTER TABLE public.end_points ALTER COLUMN id DROP DEFAULT;
       public          bushubdegefu    false    227    228    228            �           2604    49242    features id    DEFAULT     j   ALTER TABLE ONLY public.features ALTER COLUMN id SET DEFAULT nextval('public.features_id_seq'::regclass);
 :   ALTER TABLE public.features ALTER COLUMN id DROP DEFAULT;
       public          bushubdegefu    false    225    226    226            �           2604    49294    jwt_salts id    DEFAULT     l   ALTER TABLE ONLY public.jwt_salts ALTER COLUMN id SET DEFAULT nextval('public.jwt_salts_id_seq'::regclass);
 ;   ALTER TABLE public.jwt_salts ALTER COLUMN id DROP DEFAULT;
       public          bushubdegefu    false    234    233    234            �           2604    49215    pages id    DEFAULT     d   ALTER TABLE ONLY public.pages ALTER COLUMN id SET DEFAULT nextval('public.pages_id_seq'::regclass);
 7   ALTER TABLE public.pages ALTER COLUMN id DROP DEFAULT;
       public          bushubdegefu    false    222    223    223            �           2604    49181    roles id    DEFAULT     d   ALTER TABLE ONLY public.roles ALTER COLUMN id SET DEFAULT nextval('public.roles_id_seq'::regclass);
 7   ALTER TABLE public.roles ALTER COLUMN id DROP DEFAULT;
       public          bushubdegefu    false    219    220    220            �           2604    49156    users id    DEFAULT     d   ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);
 7   ALTER TABLE public.users ALTER COLUMN id DROP DEFAULT;
       public          bushubdegefu    false    216    215    216            \          0    49166    apps 
   TABLE DATA                 public          bushubdegefu    false    218   5^       h          0    49272    blob_pictures 
   TABLE DATA                 public          bushubdegefu    false    230   M_       j          0    49281    blob_videos 
   TABLE DATA                 public          bushubdegefu    false    232   g_       f          0    49256 
   end_points 
   TABLE DATA                 public          bushubdegefu    false    228   �_       d          0    49239    features 
   TABLE DATA                 public          bushubdegefu    false    226   d       l          0    49291 	   jwt_salts 
   TABLE DATA                 public          bushubdegefu    false    234   fe       b          0    49223 
   page_roles 
   TABLE DATA                 public          bushubdegefu    false    224   �e       a          0    49212    pages 
   TABLE DATA                 public          bushubdegefu    false    223   �f       ^          0    49178    roles 
   TABLE DATA                 public          bushubdegefu    false    220   tg       _          0    49196 
   user_roles 
   TABLE DATA                 public          bushubdegefu    false    221   �h       Z          0    49153    users 
   TABLE DATA                 public          bushubdegefu    false    216   i       }           0    0    apps_id_seq    SEQUENCE SET     :   SELECT pg_catalog.setval('public.apps_id_seq', 1, false);
          public          bushubdegefu    false    217            ~           0    0    blob_pictures_id_seq    SEQUENCE SET     C   SELECT pg_catalog.setval('public.blob_pictures_id_seq', 1, false);
          public          bushubdegefu    false    229                       0    0    blob_videos_id_seq    SEQUENCE SET     A   SELECT pg_catalog.setval('public.blob_videos_id_seq', 1, false);
          public          bushubdegefu    false    231            �           0    0    end_points_id_seq    SEQUENCE SET     B   SELECT pg_catalog.setval('public.end_points_id_seq', 4760, true);
          public          bushubdegefu    false    227            �           0    0    features_id_seq    SEQUENCE SET     >   SELECT pg_catalog.setval('public.features_id_seq', 1, false);
          public          bushubdegefu    false    225            �           0    0    jwt_salts_id_seq    SEQUENCE SET     >   SELECT pg_catalog.setval('public.jwt_salts_id_seq', 1, true);
          public          bushubdegefu    false    233            �           0    0    pages_id_seq    SEQUENCE SET     ;   SELECT pg_catalog.setval('public.pages_id_seq', 1, false);
          public          bushubdegefu    false    222            �           0    0    roles_id_seq    SEQUENCE SET     ;   SELECT pg_catalog.setval('public.roles_id_seq', 1, false);
          public          bushubdegefu    false    219            �           0    0    users_id_seq    SEQUENCE SET     :   SELECT pg_catalog.setval('public.users_id_seq', 7, true);
          public          bushubdegefu    false    215            �           2606    49176    apps apps_name_key 
   CONSTRAINT     M   ALTER TABLE ONLY public.apps
    ADD CONSTRAINT apps_name_key UNIQUE (name);
 <   ALTER TABLE ONLY public.apps DROP CONSTRAINT apps_name_key;
       public            bushubdegefu    false    218            �           2606    49174    apps apps_pkey 
   CONSTRAINT     L   ALTER TABLE ONLY public.apps
    ADD CONSTRAINT apps_pkey PRIMARY KEY (id);
 8   ALTER TABLE ONLY public.apps DROP CONSTRAINT apps_pkey;
       public            bushubdegefu    false    218            �           2606    49279     blob_pictures blob_pictures_pkey 
   CONSTRAINT     ^   ALTER TABLE ONLY public.blob_pictures
    ADD CONSTRAINT blob_pictures_pkey PRIMARY KEY (id);
 J   ALTER TABLE ONLY public.blob_pictures DROP CONSTRAINT blob_pictures_pkey;
       public            bushubdegefu    false    230            �           2606    49288    blob_videos blob_videos_pkey 
   CONSTRAINT     Z   ALTER TABLE ONLY public.blob_videos
    ADD CONSTRAINT blob_videos_pkey PRIMARY KEY (id);
 F   ALTER TABLE ONLY public.blob_videos DROP CONSTRAINT blob_videos_pkey;
       public            bushubdegefu    false    232            �           2606    49265    end_points end_points_name_key 
   CONSTRAINT     Y   ALTER TABLE ONLY public.end_points
    ADD CONSTRAINT end_points_name_key UNIQUE (name);
 H   ALTER TABLE ONLY public.end_points DROP CONSTRAINT end_points_name_key;
       public            bushubdegefu    false    228            �           2606    49263    end_points end_points_pkey 
   CONSTRAINT     X   ALTER TABLE ONLY public.end_points
    ADD CONSTRAINT end_points_pkey PRIMARY KEY (id);
 D   ALTER TABLE ONLY public.end_points DROP CONSTRAINT end_points_pkey;
       public            bushubdegefu    false    228            �           2606    49249    features features_name_key 
   CONSTRAINT     U   ALTER TABLE ONLY public.features
    ADD CONSTRAINT features_name_key UNIQUE (name);
 D   ALTER TABLE ONLY public.features DROP CONSTRAINT features_name_key;
       public            bushubdegefu    false    226            �           2606    49247    features features_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY public.features
    ADD CONSTRAINT features_pkey PRIMARY KEY (id);
 @   ALTER TABLE ONLY public.features DROP CONSTRAINT features_pkey;
       public            bushubdegefu    false    226            �           2606    49298    jwt_salts jwt_salts_pkey 
   CONSTRAINT     V   ALTER TABLE ONLY public.jwt_salts
    ADD CONSTRAINT jwt_salts_pkey PRIMARY KEY (id);
 B   ALTER TABLE ONLY public.jwt_salts DROP CONSTRAINT jwt_salts_pkey;
       public            bushubdegefu    false    234            �           2606    49227    page_roles page_roles_pkey 
   CONSTRAINT     f   ALTER TABLE ONLY public.page_roles
    ADD CONSTRAINT page_roles_pkey PRIMARY KEY (page_id, role_id);
 D   ALTER TABLE ONLY public.page_roles DROP CONSTRAINT page_roles_pkey;
       public            bushubdegefu    false    224    224            �           2606    49222    pages pages_name_key 
   CONSTRAINT     O   ALTER TABLE ONLY public.pages
    ADD CONSTRAINT pages_name_key UNIQUE (name);
 >   ALTER TABLE ONLY public.pages DROP CONSTRAINT pages_name_key;
       public            bushubdegefu    false    223            �           2606    49220    pages pages_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.pages
    ADD CONSTRAINT pages_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.pages DROP CONSTRAINT pages_pkey;
       public            bushubdegefu    false    223            �           2606    49190    roles roles_description_key 
   CONSTRAINT     ]   ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_description_key UNIQUE (description);
 E   ALTER TABLE ONLY public.roles DROP CONSTRAINT roles_description_key;
       public            bushubdegefu    false    220            �           2606    49188    roles roles_name_key 
   CONSTRAINT     O   ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_name_key UNIQUE (name);
 >   ALTER TABLE ONLY public.roles DROP CONSTRAINT roles_name_key;
       public            bushubdegefu    false    220            �           2606    49186    roles roles_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.roles DROP CONSTRAINT roles_pkey;
       public            bushubdegefu    false    220            �           2606    49200    user_roles user_roles_pkey 
   CONSTRAINT     f   ALTER TABLE ONLY public.user_roles
    ADD CONSTRAINT user_roles_pkey PRIMARY KEY (role_id, user_id);
 D   ALTER TABLE ONLY public.user_roles DROP CONSTRAINT user_roles_pkey;
       public            bushubdegefu    false    221    221            �           2606    49164    users users_email_key 
   CONSTRAINT     Q   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);
 ?   ALTER TABLE ONLY public.users DROP CONSTRAINT users_email_key;
       public            bushubdegefu    false    216            �           2606    49162    users users_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public            bushubdegefu    false    216            �           1259    49289    idx_blob_videos_name    INDEX     L   CREATE INDEX idx_blob_videos_name ON public.blob_videos USING btree (name);
 (   DROP INDEX public.idx_blob_videos_name;
       public            bushubdegefu    false    232            �           2606    49191    roles fk_apps_roles    FK CONSTRAINT     p   ALTER TABLE ONLY public.roles
    ADD CONSTRAINT fk_apps_roles FOREIGN KEY (app_id) REFERENCES public.apps(id);
 =   ALTER TABLE ONLY public.roles DROP CONSTRAINT fk_apps_roles;
       public          bushubdegefu    false    218    220    3237            �           2606    49266     end_points fk_features_endpoints    FK CONSTRAINT     �   ALTER TABLE ONLY public.end_points
    ADD CONSTRAINT fk_features_endpoints FOREIGN KEY (feature_id) REFERENCES public.features(id);
 J   ALTER TABLE ONLY public.end_points DROP CONSTRAINT fk_features_endpoints;
       public          bushubdegefu    false    228    3255    226            �           2606    49228    page_roles fk_page_roles_page    FK CONSTRAINT     �   ALTER TABLE ONLY public.page_roles
    ADD CONSTRAINT fk_page_roles_page FOREIGN KEY (page_id) REFERENCES public.pages(id) ON UPDATE CASCADE ON DELETE CASCADE;
 G   ALTER TABLE ONLY public.page_roles DROP CONSTRAINT fk_page_roles_page;
       public          bushubdegefu    false    3249    224    223            �           2606    49233    page_roles fk_page_roles_role    FK CONSTRAINT     �   ALTER TABLE ONLY public.page_roles
    ADD CONSTRAINT fk_page_roles_role FOREIGN KEY (role_id) REFERENCES public.roles(id) ON UPDATE CASCADE ON DELETE CASCADE;
 G   ALTER TABLE ONLY public.page_roles DROP CONSTRAINT fk_page_roles_role;
       public          bushubdegefu    false    224    220    3243            �           2606    49250    features fk_roles_features    FK CONSTRAINT     �   ALTER TABLE ONLY public.features
    ADD CONSTRAINT fk_roles_features FOREIGN KEY (role_id) REFERENCES public.roles(id) ON UPDATE CASCADE;
 D   ALTER TABLE ONLY public.features DROP CONSTRAINT fk_roles_features;
       public          bushubdegefu    false    226    3243    220            �           2606    49201    user_roles fk_user_roles_role    FK CONSTRAINT     �   ALTER TABLE ONLY public.user_roles
    ADD CONSTRAINT fk_user_roles_role FOREIGN KEY (role_id) REFERENCES public.roles(id) ON UPDATE CASCADE;
 G   ALTER TABLE ONLY public.user_roles DROP CONSTRAINT fk_user_roles_role;
       public          bushubdegefu    false    3243    221    220            �           2606    49206    user_roles fk_user_roles_user    FK CONSTRAINT     �   ALTER TABLE ONLY public.user_roles
    ADD CONSTRAINT fk_user_roles_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE;
 G   ALTER TABLE ONLY public.user_roles DROP CONSTRAINT fk_user_roles_user;
       public          bushubdegefu    false    216    221    3233            \     x�e��j�0�u��sQ�lّ�*I�ڐX`;ݏ�I1���_�Z�f��{�dyy**��Jø���/8�3|���S�==G[��Jiq�]TQ�7Ȕ41��2�Ԓat�Qdkn�[��\;];졠yX'Kp����~�7[g��8tw;�A(oș�@0�b�Ў��PY*%̯�FM���T��O�}_wM��x�ʰe��)!�;<��;i��e��Z��T�uv/~�f^&\���yh�{�7�s8����+xӐ�*��W��k�e*      h   
   x���          j   
   x���          f   z  x���Ko�8�ϛO�[�ER[�Rz*Ro�����WC�TG�c	���ǯH��C��E������~\��=���Q{~>�Ň�X�ڦ>���O�����_�f��������iwj�}���U9X�eSt�՟_��K/��X�O����5�!Ź���k՟�3�<��!C� Nn6Qa/U�cwh���hy[�ߒ�vi��xA������	md���&��r���?TA���.%�]�t𖧦���<rJ6IS���@��#��z�	�ǝ�|�^�����팅&� e�)�C��=�%�)����b�6ʩ�N*9�m둔�!�'d.t�b����)�X�vQ	αk�2
�GźkY�� ��Q��|P�������)�������1�7��-�`���3a�.K]�!��
R����f�du��|���A�7��L�̑��B�R���P���(o�6�e�kNρĆ�'Ӝt�p&�X��s�$�hECGH^93�H�$ʼ{yn�'-rp�H��<G�k��aу�m�A,�l�$�S�B R�Q�׍I��:)	�H#���� ��/!X���|�&wAeA&S]�lS��3E��g*�p�<��q�:q����,�$hDE��\�g�h������ӡKO!H��DdY)��MK]�Ny�׼>�0e� �$��2u����uY<�r>����gS���e�i���U�ԕϋ�~���in�9Ԕ�Cy2�8�6~#�l���s������A����E$\�s1*<����x*
�2��+�ؖա��6���j�]^��k2�
�@��e�g��#lQ2��(s�q��{(�P���$���������)@oD;�C���RySD�)P�A�����b*��ѱL��uԙ[h��x���b�25,�P^�b�^�bo�a�w���L
�Qp�zb@��zn�x�J�ۧ��WO�hEhF�M������.x�Ԅ�6>[Y�+�:p'�D�(;T�8�&�A��Ul�� ��4B9HH�1�SɋA����ttJH.S4���YH@���=��΂z�?������#iG�孩��K~Tu���瑩��?��i���6���n�F���~��g�p�}ތ֛�ׇ���� 0��.      d   K  x���Mk�@���W�-Hi>��S�G+����6�h�]v���;�����Q�}�y���]^���s&��-48.��z�<<�\h��I#��8
,�ƂL`G���V��x�1��+d�ŗ�0E��3��.Ih	��X�a�R �nr�MHP����LoD?f$P,�c[J��7��/$`Ju�_(��Z����F6��ވǜ+)r��`]�����h?�U�Xiw�M}��2�;h�*S�f&S�W�f`/�B�r8#�8Fc�C)�da��))xy���誻�Z��e�2A_`�)5U��e���� �a��e��0��w�q�Y}0�      l   p   x���v
Q���W((M��L��*/�/N�))Vs�	u���0�QPw���)rq4O/�tLtU
��f���Z�'T�k*��)8����x:�(��+���xx��[sqq �.�      b   �   x�M�A�0��9*��sO27-�\�UT�����<���$�I]�9	��ϭ{ޗ���^ޯ�輫��V��є,�D����X$#��41`���d=q+E�"ɐd��?���X��E^(�RV.��$>��J�/h60�      a   �   x�]��N�0��5y���"$� VUIh��FMRց�RKŶ^�I0��r��/�R���A)	;��ǅ�z�]Vm^'g����J����\d�^�6�t�kޮ�ژ�ɶU􍉆�↭�E�D��e{�E�D��c+�󣋣���g����}�q(ؗ�FK��j�k�'M��e(��8}o{�}�E���Q�O�v�R`%EQ���B6�R�<%I��ds3      ^     x�m��n�0���S̍VB��ߥ��@���ke�-��ؖ����k�*�qwg>�lYmf�e�԰a�s��LO�|��m����#,��ɍ�P'B޶�=� �{8�&��k�1���Q+.���q�<%�(�)���Y��B�$Dޕ�9:T7�f/N��I�9{�y�`���=n0/	c�g�x!X��v]@��g��kT����`B����a�">�U��@���?}	�:��
���3���H��@�ƿ��_$�� �Y�Z�fB]aZW�e9mPԨ�fQV�Y�� Z���      _   g   x���v
Q���W((M��L�+-N-�/��I-Vs�	u���0�Q0�ԁ0L`S0�HG��0�1��ccM?g?7O�?�O?wk... kTB      Z     x�͔�n7���S�.-]��(R��i�����^������gѷ/�����5Ï��>}��r;?\^l_��j����n��׏�t�Ï���F;�YR�*�0h��G*F�|A�o����x����}=��|�W׻��x:m�x9H/��)�����@ǆs��6m��j�,]�.q�!��[^Rs-(�+	�R색u�G͵���M�T���&���9�nX���9�KT�'�կ����A]�Z�d�i:B*�R�18I�21"<~�x5�����8R~Μ'S�ὃp3��ŋ+
���:M�e��Ei�
N�BC̽"�jN1��q˧���;��5�*v�}8u�]�[f�aϴ�S~�,Q�Zِ�z����RW�Բ��݇����9r��cJTL��VMsV\�!����3��}*2&Ƒ�nLS��g�FX- �����$:�e����$9Iz�-|��y_tX��$ƓF:��
	0k8VҦ�,�S'�z{����W�=�ƻ��jLX7�=ɞyWI	�K^��0nh9[baN�����#���뚏���f������p�{��R���9���ѥqIE�R����i�aKb��w�W7��7�7�{�yG�������.�R�G�-�^�4�ԣ�$��
�"�
9����o����V�=�ՠ�n����V�<r�Uɲ:';�ˊ�W��2Uk⣕5��5^X�F�E:�E�jpF�b	��>�����1�g�n���������/��.�����?~9;;��K��     