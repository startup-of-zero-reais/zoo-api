CREATE TYPE age_enum AS ENUM ('neonate','cub','young','adult','senile');

CREATE TYPE gender_enum AS ENUM ('male','female','undefined');


CREATE TABLE animals (
  id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
  name VARCHAR(255) DEFAULT NULL,
  washer_code VARCHAR(255) NOT NULL,
  microchip_code VARCHAR(255) NOT NULL,
  landing_at TIMESTAMP NOT NULL,
  origin VARCHAR(255) NOT NULL,
  born_date TIMESTAMP DEFAULT NULL,
  age age_enum DEFAULT NULL,
  gender gender_enum DEFAULT NULL,
  observation VARCHAR(255) DEFAULT NULL,
  species_id UUID NOT NULL,
  enclosure_id UUID NOT NULL,
  

  search_vector TSVECTOR GENERATED ALWAYS AS (
    setweight(to_tsvector('portuguese', name), 'A') ||
    setweight(to_tsvector('portuguese', washer_code), 'B') ||
    setweight(to_tsvector('portuguese', microchip_code), 'C')
  ) STORED,
 
  CONSTRAINT fk_species_id FOREIGN KEY(species_id) REFERENCES species(id) ON DELETE CASCADE,
  CONSTRAINT fk_enclosure_id FOREIGN KEY(enclosure_id) REFERENCES enclosures(id) ON DELETE CASCADE,

  UNIQUE (id)
) INHERITS (_timestamps);


CREATE INDEX IF NOT EXISTS idx_animals_search_vector ON animals USING GIN(search_vector);
CREATE INDEX IF NOT EXISTS deleted_at_idx ON animals(deleted_at) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_animals_updated_at ON animals(updated_at DESC);
CREATE INDEX IF NOT EXISTS idx_washer_code ON animals(washer_code);
CREATE INDEX IF NOT EXISTS idx_microchip_code ON animals(microchip_code);

CREATE OR REPLACE TRIGGER set_updated_at_species
BEFORE UPDATE ON animals FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();


