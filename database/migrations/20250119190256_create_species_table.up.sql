CREATE TABLE species (
  id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
  common_name VARCHAR(255) NOT NULL,
  scientific_name VARCHAR(255) NOT NULL,
  taxonomic_order VARCHAR(255) NOT NULL,
  kind VARCHAR(255) NOT NULL,

  search_vector TSVECTOR GENERATED ALWAYS AS (
    setweight(to_tsvector('portuguese', common_name), 'A') ||
    setweight(to_tsvector('portuguese', scientific_name), 'B') ||
    setweight(to_tsvector('portuguese', kind), 'C') ||
    setweight(to_tsvector('portuguese', taxonomic_order), 'D')
  ) STORED,

  UNIQUE (id)
) INHERITS (_timestamps);


CREATE INDEX IF NOT EXISTS idx_species_search_vector ON species USING GIN(search_vector);
CREATE INDEX IF NOT EXISTS deleted_at_idx ON species(deleted_at) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_species_updated_at ON species(updated_at DESC);

CREATE OR REPLACE TRIGGER set_updated_at_species
BEFORE UPDATE ON species FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
