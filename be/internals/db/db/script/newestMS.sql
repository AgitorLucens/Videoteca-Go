CREATE OR REPLACE FUNCTION ObtenerUltimas10Peliculas()
RETURNS SETOF movie_series
LANGUAGE plpgsql
AS $BODY$
BEGIN
    
    RETURN QUERY 
    SELECT 
        ms_id, 
        title, 
        synopsis, 
        release_year, 
        classification_ms, 
        director, 
        cover, 
        date_added, 
        duration, 
        ms_type, 
        trailer 
    FROM 
        movie_series
    ORDER BY 
        ms_id DESC
    LIMIT 10; 
    
END;
$BODY$