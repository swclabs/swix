CREATE OR REPLACE FUNCTION update_product_rating()
RETURNS TRIGGER AS $$
BEGIN
  IF OLD.rating <> -1.0 AND NEW.rating IS DISTINCT FROM OLD.rating THEN
    NEW.rating := (OLD.rating + NEW.rating) / 2;
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER recalculate_rating
BEFORE UPDATE ON products
FOR EACH ROW
EXECUTE FUNCTION update_product_rating();