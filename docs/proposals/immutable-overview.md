# Immutable

One attribute per object. Its effect depends on the object type:
on a file it means **freeze**, on a directory it means **protect**.

## File — freeze

The file is final:

- No modification (no new revision).
- No deletion.
- No renaming / moving.

Since every change creates a new revision, "no new revision" also means
"content is fixed".

## Directory — protect

A directory is a container, defined by its entries. Protecting it fixes
the container, not the entries:

- No entries can be added or removed.
- No modification of direct entries (files or subdirectories).
- No renaming / moving / deleting of the directory itself.

It does **not** propagate to the entries: a protected directory can
contain an unprotected directory in which entries can still be added,
removed or modified.

## Self vs. parent

An object is immutable if its own attribute is set (self) **or** its
parent directory's attribute is set (parent). Renaming or deleting an
object is governed by the parent; adding or removing inside it by the
object itself.

## Setting

- Single: this object only.
- Recursive: this object and all descendants (a batch operation; one
  attribute is still stored per object).
