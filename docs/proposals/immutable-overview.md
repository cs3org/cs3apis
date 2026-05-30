# Immutable

One attribute per object. Its effect depends on the object type:
on a file it means **freeze**, on a directory it means **protect**.

## File — freeze

The file is final:

- No modification (content is fixed).
- No deletion.
- No renaming / moving.
- **Irreversible** — once set, the attribute cannot be removed.

## Directory — protect

A directory is a container, defined by its entries. Protecting it fixes
the container, not the entries:

- No entries can be added or removed.
- No modification of direct entries (files or subdirectories).
- No renaming / moving / deleting of the directory itself.
- **Reversible** — space managers or administrators can remove the attribute.

It does **not** propagate to the entries: a protected directory can
contain an unprotected directory in which entries can still be added,
removed or modified.

## Self vs. parent

An object is immutable if its own attribute is set (self) **or** its
parent directory's attribute is set (parent). Renaming, modification
or deleting an object is governed by the parent; adding or removing
inside it by the object itself.

## Effective state

Each object has one of three effective states:

| State | Meaning | Icon |
|-------|---------|------|
| **Frozen** | Own attribute is set | Shield filled |
| **Protected** | Parent's attribute is set | Shield outline |
| **None** | Neither self nor parent | — |

This applies equally to files and directories. A directory inside a
protected parent is itself Protected — it cannot be deleted or renamed,
but since it is not itself immutable, entries inside it can be freely
added, removed or modified.

## Distinction from locks

| | Lock | Immutable |
|---|---|---|
| Purpose | Collaborative editing | Structure/content protection |
| Duration | Temporary (with expiration) | Permanent (files) / until admin removes (dirs) |
| Set by | Any user with write access | Manager / Admin only |
| Scope | Prevents parallel writes | Prevents all modifications |

## Permissions

Setting the immutable attribute requires separate permissions for
files and directories:

- `set_immutable_file` — freeze files (irreversible, grant carefully)
- `set_immutable_container` — protect/unprotect containers (reversible)

Only Manager and Coowner roles have these permissions by default.

## Setting

- **Single**: `SetImmutable` RPC on a single resource reference.
- **Recursive**: application-level batch operation (e.g. `aktenplan-apply`
  calling `SetImmutable` per node). One attribute is stored per object.
  Recursive setting is not part of the CS3 API — it is the caller's
  responsibility.

## Storage

Persistent as extended attribute `user.oc.immutable` on the filesystem
node. Returned in `ResourceInfo.Immutable` via `Stat()`. Effective
state (frozen/protected) transported via `Opaque["immutable-state"]`.

## WebDAV

- `oc:immutable` property: `frozen`, `protected`, or absent
- `oc:permissions`: `D` (delete) and `NV` (rename/move) flags are
  stripped when the resource is effectively immutable

## Implementation status

Fully implemented across all layers:

| Layer | Status |
|-------|--------|
| CS3 Proto (cs3apis#272) | APPROVED, awaiting merge |
| Go Bindings (go-cs3apis fork) | Done |
| Reva decomposedfs | Done, tested (21+ tests) |
| WebDAV PROPFIND | Done |
| Graph API (OpenCloud) | Done |
| Web Frontend | Done |
