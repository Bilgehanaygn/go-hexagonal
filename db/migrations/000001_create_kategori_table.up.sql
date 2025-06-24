CREATE TABLE kategori (
    created_date TIMESTAMP WITHOUT TIME ZONE DEFAULT now(),
    last_modified_date TIMESTAMP WITHOUT TIME ZONE DEFAULT now(),
    version INTEGER DEFAULT 0,
    isim VARCHAR(50) NOT NULL,
    tanimlayici_deger VARCHAR(50) NOT NULL,
    tur VARCHAR(32) NOT NULL,
    durum VARCHAR(10),
    gerekce VARCHAR(255),
    id UUID NOT NULL,
    ebeveyn_kategori_id UUID,
    sira_no BIGINT,
    CONSTRAINT kategori_pkey PRIMARY KEY (id)
);