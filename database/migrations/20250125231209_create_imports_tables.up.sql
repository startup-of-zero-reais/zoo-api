CREATE TYPE import_status AS ENUM ('sending', 'received', 'processing', 'completed', 'error');

CREATE TABLE IF NOT EXISTS import_state (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    filename TEXT NOT NULL,
    upload_id TEXT NOT NULL,
    state import_status NOT NULL DEFAULT 'sending',

    UNIQUE (id),
    UNIQUE (upload_id)
);

CREATE TABLE IF NOT EXISTS import_enclosures (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    identification VARCHAR(255) DEFAULT NULL,
    reason VARCHAR(255) DEFAULT NULL,
    state_id UUID NOT NULL,

    UNIQUE (id),
    UNIQUE (identification),

    CONSTRAINT fk_enclosure_state_id FOREIGN KEY(state_id) REFERENCES import_state(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS import_species (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    common_name VARCHAR(255) DEFAULT NULL,
    scientific_name VARCHAR(255) DEFAULT NULL,
    taxonomic_order VARCHAR(255) DEFAULT NULL,
    kind VARCHAR(255) DEFAULT NULL,
    reason VARCHAR(255) DEFAULT NULL,
    state_id UUID NOT NULL,

    UNIQUE (id),
    UNIQUE (common_name),

    CONSTRAINT fk_species_state_id FOREIGN KEY(state_id) REFERENCES import_state(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS import_animals (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(255) DEFAULT NULL,
    washer_code VARCHAR(255) DEFAULT NULL,
    microchip_code VARCHAR(255) DEFAULT NULL,
    landing_at TIMESTAMP DEFAULT NULL,
    origin VARCHAR(255) DEFAULT NULL,
    born_date TIMESTAMP DEFAULT NULL,
    age VARCHAR(30) DEFAULT NULL,
    gender VARCHAR(50) DEFAULT NULL,
    observation VARCHAR(255) DEFAULT NULL,
    reason VARCHAR(255) DEFAULT NULL,
    state_id UUID NOT NULL,

    species_id UUID DEFAULT NULL,
    enclosure_id UUID DEFAULT NULL,

    UNIQUE (id),
    UNIQUE (microchip_code),
    UNIQUE (washer_code),

    CONSTRAINT fk_animals_state_id FOREIGN KEY(state_id) REFERENCES import_state(id) ON DELETE CASCADE
);


