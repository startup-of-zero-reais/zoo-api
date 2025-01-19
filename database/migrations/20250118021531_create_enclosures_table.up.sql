CREATE TABLE enclosures (
  id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
  identification VARCHAR(255) NOT NULL,
  search_vector TSVECTOR GENERATED ALWAYS AS (to_tsvector('portuguese', identification)) STORED,

  UNIQUE (id)
) INHERITS (_timestamps);


CREATE INDEX IF NOT EXISTS idx_enclsoure_search_vector ON enclosures USING GIN(search_vector);
CREATE INDEX IF NOT EXISTS deletet_at_idx ON enclosures(deleted_at) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_enclosures_updated_at ON enclosures(updated_at DESC);

CREATE OR REPLACE TRIGGER set_updated_at_enclosures
BEFORE UPDATE ON enclosures FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
