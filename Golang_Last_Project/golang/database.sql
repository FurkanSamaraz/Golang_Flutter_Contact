PGDMP     #    :                z            postgres    14.1    14.1     ?           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            ?           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            ?           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            ?           1262    14020    postgres    DATABASE     S   CREATE DATABASE postgres WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'C';
    DROP DATABASE postgres;
                postgres    false            ?           0    0    DATABASE postgres    COMMENT     N   COMMENT ON DATABASE postgres IS 'default administrative connection database';
                   postgres    false    3580                        3079    16384 	   adminpack 	   EXTENSION     A   CREATE EXTENSION IF NOT EXISTS adminpack WITH SCHEMA pg_catalog;
    DROP EXTENSION adminpack;
                   false            ?           0    0    EXTENSION adminpack    COMMENT     M   COMMENT ON EXTENSION adminpack IS 'administrative functions for PostgreSQL';
                        false    2            ?            1259    16456    userr    TABLE     ?   CREATE TABLE public.userr (
    id integer NOT NULL,
    username character varying(100) NOT NULL,
    password character varying(100) NOT NULL,
    eposta character varying(100) NOT NULL
);
    DROP TABLE public.userr;
       public         heap    postgres    false            ?            1259    16455    userr_id_seq    SEQUENCE     ?   CREATE SEQUENCE public.userr_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.userr_id_seq;
       public          postgres    false    211            ?           0    0    userr_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.userr_id_seq OWNED BY public.userr.id;
          public          postgres    false    210            g           2604    16459    userr id    DEFAULT     d   ALTER TABLE ONLY public.userr ALTER COLUMN id SET DEFAULT nextval('public.userr_id_seq'::regclass);
 7   ALTER TABLE public.userr ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    210    211    211            ?          0    16456    userr 
   TABLE DATA           ?   COPY public.userr (id, username, password, eposta) FROM stdin;
    public          postgres    false    211   ?                   0    0    userr_id_seq    SEQUENCE SET     ;   SELECT pg_catalog.setval('public.userr_id_seq', 58, true);
          public          postgres    false    210            i           2606    16461    userr userr_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.userr
    ADD CONSTRAINT userr_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.userr DROP CONSTRAINT userr_pkey;
       public            postgres    false    211            ?   ?  x????N1@??W??nv??ۃ??L?U"??D???????e?v?	?Ü3:? L???7	n???J.???f5@?&<,????K=,?0}|?7??|??t?7??R?˗?U	971??|?5S=?']??}ʏ?/??!?`R??[??????U?!?$??Mj)???r?l]4::??dn?Dm	?j?n.秺f~??O??S?Ά??yjr?{?:Uu?????ߟ????/??????d??]??ryغ???$!??????S?}U?????2{?\???[??uQ??[???끨-?ipbפ?E?-f?U?o?Ȫ\y??Ƣ?g?+?K?:4??a?????(
????@?u$???!?]??3????m??c5w????F?n?
+c3u7?u`????Y?@???kp?Tu?rȢ$?*?U?
c??tv?fD???iqPM?v??崕???8??r?(     