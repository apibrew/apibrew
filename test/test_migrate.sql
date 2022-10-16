-- create database dh_test;
-- create user dh_test superuser login password 'dh_test';

create table if not exists public.organization (
                                     organization_uuid uuid primary key not null,
                                     founded_on date,
                                     founded_on_precision character varying(12),
                                     legal_name character varying(1024),
                                     operating_status character varying(32),
                                     short_description text,
                                     title character varying(1024),
                                     website character varying(2048),
                                     ipo_status character varying(32),
                                     num_employees_enum character varying(32),
                                     last_funding_type character varying(32),
                                     rank_org_company character varying(32)
);
create index organization_founded_on_idx on organization using btree (founded_on);


INSERT INTO public.organization (organization_uuid) VALUES ('74994fdd-fc2f-b98b-76a0-48862cb6444d');
INSERT INTO public.organization (organization_uuid) VALUES ('e361535b-948b-43d6-bdde-cfd589a8c9ce');
INSERT INTO public.organization (organization_uuid) VALUES ('0af9222b-fe3f-b65e-ffcc-60e2c3b6bb3d');
INSERT INTO public.organization (organization_uuid) VALUES ('1c979dd7-56cf-0e06-9115-7870812ed329');
INSERT INTO public.organization (organization_uuid) VALUES ('c77f5bdb-1e08-e875-9cbd-e72133dda6df');
INSERT INTO public.organization (organization_uuid) VALUES ('ad633b25-46e8-49e5-bb57-499e4494df71');
INSERT INTO public.organization (organization_uuid) VALUES ('df9080d2-4cb1-4018-9727-a24e13c9c64e');
INSERT INTO public.organization (organization_uuid) VALUES ('74f3612b-046d-4831-9182-da68ee90e83b');
INSERT INTO public.organization (organization_uuid) VALUES ('48a17caf-57c7-81b5-0818-9af7c16d29db');
INSERT INTO public.organization (organization_uuid) VALUES ('68bdf7d1-b1cd-4b42-a4da-6d9c6723caa1');
INSERT INTO public.organization (organization_uuid) VALUES ('fd2f9cf7-eb8e-384b-7c85-83958baba54c');
INSERT INTO public.organization (organization_uuid) VALUES ('ab251475-55a5-7d0e-153c-ec1e87f78f0c');
INSERT INTO public.organization (organization_uuid) VALUES ('e0193fa5-19c1-8008-ccae-ead1c67fadd6');
INSERT INTO public.organization (organization_uuid) VALUES ('da04be2d-41c6-4a80-9e75-2aeee99dc096');
INSERT INTO public.organization (organization_uuid) VALUES ('4c459db5-a825-a3fe-d59d-fdc443efd225');
INSERT INTO public.organization (organization_uuid) VALUES ('03f82643-f498-4b5e-9f61-f9b18ff5b9ee');
INSERT INTO public.organization (organization_uuid) VALUES ('62950e72-ff64-54aa-cbe0-9d5cd2bb8102');
INSERT INTO public.organization (organization_uuid) VALUES ('b80cf72e-4c9f-406a-b64a-2c785987e8b9');
INSERT INTO public.organization (organization_uuid) VALUES ('176c1306-58b8-4c51-af8c-f87cfa21b43f');
INSERT INTO public.organization (organization_uuid) VALUES ('ff29a974-1d89-4908-8b6a-99732556119a');
INSERT INTO public.organization (organization_uuid) VALUES ('bce0b489-aad7-46a5-9ec3-838870448755');
INSERT INTO public.organization (organization_uuid) VALUES ('f9b1cddc-c817-3cb5-fe27-78b88b7edcee');
INSERT INTO public.organization (organization_uuid) VALUES ('8e7fcad5-5ef8-be9c-2cee-f2c6cbf533f9');
INSERT INTO public.organization (organization_uuid) VALUES ('95f746e4-4c75-e125-39c3-6eb92847c7c4');
INSERT INTO public.organization (organization_uuid) VALUES ('a13cde81-38cb-8caf-ebbb-4c9cbbc6a773');
INSERT INTO public.organization (organization_uuid) VALUES ('8d9971f9-937b-7469-1446-b7fba7b30f5f');
INSERT INTO public.organization (organization_uuid) VALUES ('1744abd2-df96-4477-83dd-48860338400d');
INSERT INTO public.organization (organization_uuid) VALUES ('5e3b61e6-c335-c616-817a-cd1bc0bc1a0b');
INSERT INTO public.organization (organization_uuid) VALUES ('6f6c10e7-d12f-41ed-8e9b-932b3523ed14');
INSERT INTO public.organization (organization_uuid) VALUES ('57c6c083-225f-4ae1-b0e5-e7330275432a');
INSERT INTO public.organization (organization_uuid) VALUES ('1c33ed2e-a0ea-2299-e0b7-cc7b6fec4fa5');
INSERT INTO public.organization (organization_uuid) VALUES ('6ac05830-799c-4998-ac19-93fcb2a2d301');
INSERT INTO public.organization (organization_uuid) VALUES ('4f78996b-51da-9602-418f-4db497ffc876');
INSERT INTO public.organization (organization_uuid) VALUES ('32729a8e-b672-6423-ee12-808639697ccf');
INSERT INTO public.organization (organization_uuid) VALUES ('7962e254-2689-25db-27ed-cb4d72dd87d8');
INSERT INTO public.organization (organization_uuid) VALUES ('ebf80f49-eec6-61ba-5e3a-897ebca86255');
INSERT INTO public.organization (organization_uuid) VALUES ('29116ac2-ef20-8adc-c89b-dd3d7e1ca722');
INSERT INTO public.organization (organization_uuid) VALUES ('45b39c72-165c-4fdb-9c24-ebf352369591');
INSERT INTO public.organization (organization_uuid) VALUES ('3de9581c-c5f9-40c0-85f3-c8783ab24b52');
INSERT INTO public.organization (organization_uuid) VALUES ('3e6e571d-846e-f456-79f4-9003af50e3b1');
INSERT INTO public.organization (organization_uuid) VALUES ('73f30535-10a7-497c-a88d-72847b9017d8');
INSERT INTO public.organization (organization_uuid) VALUES ('d19a5a97-6921-4548-b94b-a46811e422ec');
INSERT INTO public.organization (organization_uuid) VALUES ('a0f6b335-4905-4d2b-8f07-65130fe0fbf4');
INSERT INTO public.organization (organization_uuid) VALUES ('dc09bcc1-db09-9bc4-75b8-2d13eb7163c9');
INSERT INTO public.organization (organization_uuid) VALUES ('cda14d27-2270-32d9-41ed-41454af09853');
INSERT INTO public.organization (organization_uuid) VALUES ('2a1a0575-7e81-4c00-a279-cde2bf3d9010');
INSERT INTO public.organization (organization_uuid) VALUES ('e2313011-1890-457e-980a-220a1f69e2ca');
INSERT INTO public.organization (organization_uuid) VALUES ('d7e6ad7b-4a3d-4d91-9d85-83e28e072978');
INSERT INTO public.organization (organization_uuid) VALUES ('d2b43c14-4fdd-baf4-0fa4-a8b28070065d');
INSERT INTO public.organization (organization_uuid) VALUES ('62046a38-551b-46b7-8435-b6005b3c4e68');
INSERT INTO public.organization (organization_uuid) VALUES ('38507b6a-d008-419e-835c-e3587c8e029e');
INSERT INTO public.organization (organization_uuid) VALUES ('348ac753-44d3-470f-9cb9-eaf43b4c983c');
INSERT INTO public.organization (organization_uuid) VALUES ('5e0fe641-f66d-4486-924c-900339c136d5');
INSERT INTO public.organization (organization_uuid) VALUES ('770c4ed2-6efa-0673-8fea-a11d78370d0b');
INSERT INTO public.organization (organization_uuid) VALUES ('993dddca-fb6e-4330-a858-ad8f2ef0097e');
INSERT INTO public.organization (organization_uuid) VALUES ('6d00421c-186d-420b-920e-ee0e4a7b5e66');
INSERT INTO public.organization (organization_uuid) VALUES ('2d36c884-cfed-4947-af08-50fefbf0b752');
INSERT INTO public.organization (organization_uuid) VALUES ('be7b9f12-c57c-a916-4989-03dbb84488b3');
INSERT INTO public.organization (organization_uuid) VALUES ('c9890f93-d7ac-4093-9e76-885495d92f1e');
INSERT INTO public.organization (organization_uuid) VALUES ('a659d10f-b94b-423b-a253-9d64d596dde8');
INSERT INTO public.organization (organization_uuid) VALUES ('33ae9d61-501a-67de-592c-fc6688a90d5e');
INSERT INTO public.organization (organization_uuid) VALUES ('3e228437-b714-487e-8249-f74c882011ec');
INSERT INTO public.organization (organization_uuid) VALUES ('89f7d7fb-ba98-41a9-ab5a-0183cd0f23ac');
INSERT INTO public.organization (organization_uuid) VALUES ('433ffb5f-cff1-47bc-9dc7-2d1a347b97e4');
INSERT INTO public.organization (organization_uuid) VALUES ('2c8a15d8-e3f1-175e-fa1f-5de8a3ce5fdf');
INSERT INTO public.organization (organization_uuid) VALUES ('b1e0dc99-cccc-4b43-830f-54168ae5208a');
INSERT INTO public.organization (organization_uuid) VALUES ('95efa9d8-5b85-f450-8d21-5cb21066b4c3');
INSERT INTO public.organization (organization_uuid) VALUES ('37abeb4c-68be-278b-b55e-ee088aae666d');
INSERT INTO public.organization (organization_uuid) VALUES ('b8d78338-813f-476f-947c-bb1194339d1d');
INSERT INTO public.organization (organization_uuid) VALUES ('e92af51d-50d5-91ba-44ea-d7529aadb839');
INSERT INTO public.organization (organization_uuid) VALUES ('fa77b0ce-bc43-c267-8d6e-9484aaeab7e5');
INSERT INTO public.organization (organization_uuid) VALUES ('bca97c67-18d2-2698-ceda-b8cbb30bcbc6');
INSERT INTO public.organization (organization_uuid) VALUES ('ae955bf3-0de4-48ea-ab63-2c03efeced6e');
INSERT INTO public.organization (organization_uuid) VALUES ('928806fe-ed50-41b7-b459-6d6798aa7b69');
INSERT INTO public.organization (organization_uuid) VALUES ('5b8644cd-6ba2-c551-3279-0998f0903a8c');
INSERT INTO public.organization (organization_uuid) VALUES ('2cf61fd7-e51d-7559-87c9-f9d0ebef9916');
INSERT INTO public.organization (organization_uuid) VALUES ('68a2005e-237b-cbec-865c-0161d83d6c09');
INSERT INTO public.organization (organization_uuid) VALUES ('b3b69a0c-84c5-40f2-b0cd-15331383c401');
INSERT INTO public.organization (organization_uuid) VALUES ('01199857-4478-4c34-a159-d7a55e56754d');
INSERT INTO public.organization (organization_uuid) VALUES ('fb91f7f3-63ec-4e7c-b8b4-a7e201fe7b96');
INSERT INTO public.organization (organization_uuid) VALUES ('d68ebe7a-c673-8505-e3ef-daffabedbabf');
INSERT INTO public.organization (organization_uuid) VALUES ('28e1bbeb-536a-47ac-84b9-73a81169092b');
INSERT INTO public.organization (organization_uuid) VALUES ('e37e3c7e-e6a1-48d5-92a7-65e9e2753271');
INSERT INTO public.organization (organization_uuid) VALUES ('7641458f-49ed-4427-b8d4-6f8bbf914c5c');
INSERT INTO public.organization (organization_uuid) VALUES ('dfdf1337-06a7-4a94-b0cd-cca577c753df');
INSERT INTO public.organization (organization_uuid) VALUES ('7169ff97-dcc7-164d-7225-60ba9c7935a0');
INSERT INTO public.organization (organization_uuid) VALUES ('515128f6-53cf-43b6-bd03-44db4f47e9eb');
INSERT INTO public.organization (organization_uuid) VALUES ('4974d65d-2af5-44d7-a315-c05f20be7070');
INSERT INTO public.organization (organization_uuid) VALUES ('4f1fc4d5-8500-4fc3-9ff3-16313c7bbc31');
INSERT INTO public.organization (organization_uuid) VALUES ('b218af47-fb17-bb77-e356-b967b8f295f4');
INSERT INTO public.organization (organization_uuid) VALUES ('7d1c7640-6d47-1e16-47c8-03163a1e7f8a');
INSERT INTO public.organization (organization_uuid) VALUES ('32e0a608-be76-49c3-82f6-80debea5df2c');
INSERT INTO public.organization (organization_uuid) VALUES ('c81bf70d-3e1d-48d7-b555-d75b2c786b32');
INSERT INTO public.organization (organization_uuid) VALUES ('953346ac-2355-4331-9817-65f2129fec39');
INSERT INTO public.organization (organization_uuid) VALUES ('2c5b05f7-9aa3-ca54-49a8-58641da46729');
INSERT INTO public.organization (organization_uuid) VALUES ('6ec8915c-9589-a4d8-963d-6ae748ee4e52');
INSERT INTO public.organization (organization_uuid) VALUES ('438b762f-432b-422d-8776-4a4e5707ef84');
INSERT INTO public.organization (organization_uuid) VALUES ('17533358-540c-c9b1-0e44-83e4b8c47466');
INSERT INTO public.organization (organization_uuid) VALUES ('b70277ac-68a2-490f-8bdf-33460919fb3d');
INSERT INTO public.organization (organization_uuid) VALUES ('1ca5b430-0f62-4151-88ea-0e7622da0c41');
INSERT INTO public.organization (organization_uuid) VALUES ('0c69c03d-701f-4599-9776-6d16ba05cef2');
INSERT INTO public.organization (organization_uuid) VALUES ('80967bac-ae02-4b37-8ee6-d63aca0bf493');
INSERT INTO public.organization (organization_uuid) VALUES ('76c80fc5-5856-455c-a6bd-f881f432fa3b');
INSERT INTO public.organization (organization_uuid) VALUES ('6ecc88f7-b40d-400d-8f65-eddc9ae4a7a9');
INSERT INTO public.organization (organization_uuid) VALUES ('2dfaa318-4b27-49ca-be62-cd3425f091ca');
INSERT INTO public.organization (organization_uuid) VALUES ('1ff42048-84af-4320-a09b-4b1f8af3dc5e');
INSERT INTO public.organization (organization_uuid) VALUES ('75649d2e-8e8a-824b-eb18-d790a4a9ba7c');
INSERT INTO public.organization (organization_uuid) VALUES ('3458d3a9-eba2-45d6-95ae-15ecb0a12f59');
INSERT INTO public.organization (organization_uuid) VALUES ('2f021422-d04e-435a-80b4-b9c7e5dba7e8');
