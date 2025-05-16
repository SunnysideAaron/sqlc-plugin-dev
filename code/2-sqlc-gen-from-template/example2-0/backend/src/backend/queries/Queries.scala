package backend.queries

import com.augustnagro.magnum
import com.augustnagro.magnum.*



case class Row_GetAuthor(
    id:bigserial,
    name:text,
    bio:Option[text],
)

def GetAuthor(
  id:bigserial,
)(using con: DbCon): Row_GetAuthor = {
  Frag("""
  SELECT id, name, bio FROM authors
WHERE id = $1 LIMIT 1
  """, params = IArray(
  id,
  )).query[Row_GetAuthor].run().head
}



case class Row_ListAuthors(
    id:bigserial,
    name:text,
    bio:Option[text],
)

def ListAuthors(
)(using con: DbCon): Vector[Row_ListAuthors] = {
  Frag("""
  SELECT id, name, bio FROM authors
ORDER BY name
  """, params = IArray()).query[Row_ListAuthors].run()
}



case class Row_CreateAuthor(
    id:bigserial,
    name:text,
    bio:Option[text],
)

def CreateAuthor(
  name:text,
  bio:text,
)(using con: DbCon): Row_CreateAuthor = {
  Frag("""
  INSERT INTO authors (
  name, bio
) VALUES (
  $1, $2
)
RETURNING id, name, bio
  """, params = IArray(
  name,
  
  bio,
  )).query[Row_CreateAuthor].run().head
}



def UpdateAuthor(
  id:bigserial,
  name:text,
  bio:text,
)(using con: DbCon): Unit = {
  Frag("""
  UPDATE authors
  set name = $2,
  bio = $3
WHERE id = $1
  """, params = IArray(
  id,
  
  name,
  
  bio,
  )).update.run()
}



def DeleteAuthor(
  id:bigserial,
)(using con: DbCon): Unit = {
  Frag("""
  DELETE FROM authors
WHERE id = $1
  """, params = IArray(
  id,
  )).update.run()
}