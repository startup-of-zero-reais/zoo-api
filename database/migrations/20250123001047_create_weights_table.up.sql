CREATE TABLE weight_history (
  id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
  animal_id UUID NOT NULL,
  user_id  UUID NOT NULL,
  weight DECIMAL(10,2) NOT NULL,

  CONSTRAINT fk_animal_id FOREIGN KEY(animal_id) REFERENCES animals(id) ON DELETE CASCADE,
  CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,

  UNIQUE (id)
) INHERITS (_timestamps);


CREATE INDEX IF NOT EXISTS deletet_at_idx ON weight_history(deleted_at) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_weight_history_updated_at ON weight_history(updated_at DESC);
CREATE INDEX IF NOT EXISTS idx_weight_history_weight ON weight_history(weight);
CREATE INDEX IF NOT EXISTS idx_weight_history_user_id ON weight_history(user_id);
CREATE INDEX IF NOT EXISTS idx_weight_history_animal_id ON weight_history(animal_id);

CREATE OR REPLACE TRIGGER set_updated_at_weights
BEFORE UPDATE ON weight_history FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
