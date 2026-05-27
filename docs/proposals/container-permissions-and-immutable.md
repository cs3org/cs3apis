# Proposal: Container-Specific Permissions and Immutable Resources

## Status

Draft

## Authors

- @flash7777

## Summary

This proposal adds three new fields to the CS3 storage provider API:

1. **`delete_container`** (ResourcePermissions field 21) — granular delete permission for containers/directories
2. **`move_container`** (ResourcePermissions field 22) — granular move/rename permission for containers/directories
3. **`immutable`** (ResourceInfo field 20) — flag to mark a resource as frozen/immutable

## Motivation

### Problem 1: No distinction between file and container operations

The current `ResourcePermissions` message uses a single `delete` flag for both files and containers, and a single `move` flag for both. This makes it impossible to configure a role that allows users to delete files but not directories, or to rename files but not directories.

This is a common requirement in Document Management Systems (DMS), records management, and compliance scenarios where directory structures must be protected while allowing normal file operations within them.

**Real-world example: File plans (Aktenplan)**

In German public administration and corporate environments, a rigid hierarchical directory structure (Aktenplan) must be maintained. Users should be able to add and edit documents within the structure, but must not be able to delete or rename the directories that form the structure.

```
/01 Administration          <- protected (no delete/rename)
  /01.01 HR                 <- protected
    /01.01.01 Recruiting    <- protected
      /Applications/        <- user can create, edit, delete files here
        resume.pdf
        cover_letter.pdf
```

### Problem 2: No immutability concept

There is no way to mark a resource as immutable/frozen in the CS3 API. Locks exist but serve a different purpose (temporary collaborative editing locks with expiration). Immutability is a permanent (until explicitly removed) state that:

- Prevents modification, deletion, moving, and renaming of the resource
- For containers: also prevents creation of new children
- Is typically set by administrators or space managers
- Supports DMS features like retention, legal hold, and archive states

## Specification

### New fields in ResourcePermissions

```protobuf
message ResourcePermissions {
  // ... existing fields 1-20 ...

  // When set, controls whether containers (directories) can be deleted,
  // independent of the delete permission which then only applies to files.
  bool delete_container = 21;

  // When set, controls whether containers (directories) can be
  // moved or renamed, independent of the move permission which
  // then only applies to files.
  bool move_container = 22;
}
```

### Backward compatibility

When `delete_container` and `move_container` are not explicitly set (default `false` in protobuf), implementations SHOULD fall back to the existing `delete` and `move` permissions for containers. This means:

- Existing clients that don't know about the new fields continue to work unchanged
- Existing roles that set `delete=true` will still allow container deletion (via fallback)
- Only when `delete_container` is explicitly managed does the separation take effect

**Recommended implementation logic:**

```
canDeleteContainer = perm.delete_container OR (perm.delete AND NOT explicitly_managing_container_perms)
```

### New field in ResourceInfo

```protobuf
message ResourceInfo {
  // ... existing fields 1-19 ...

  // When true, the resource is immutable (frozen).
  bool immutable = 20;
}
```

### Immutable semantics

For **files:**
- Cannot be modified (upload/overwrite denied)
- Cannot be deleted or moved/renamed

For **containers:**
- Cannot be deleted or moved/renamed
- No new children can be created within
- Existing non-immutable children can still be modified
- Existing immutable children follow their own immutable rules

### Setting immutable

The `immutable` flag can be set via:
- `SetArbitraryMetadata` with reserved key `cs3:immutable`
- Or a dedicated RPC (future extension)

Only users with management permissions (space owner, space manager, administrator) should be able to set or clear the immutable flag.

## ACL representation

For storage backends using ACL strings (e.g., EOS-style):

```
+dc  = delete_container allowed
!dc  = delete_container denied
+mc  = move_container allowed
!mc  = move_container denied
```

## Use cases

1. **File plan (Aktenplan) protection**: Rigid directory structure with free file operations below
2. **Retention / Legal hold**: Mark resources as immutable during retention periods
3. **Archive directories**: Freeze completed project folders while keeping them accessible
4. **Collaborative workspaces**: Allow file operations while preventing accidental directory deletion
5. **Compliance**: Meet regulatory requirements for records management

## Impact

- **Proto changes**: 3 new fields (backward compatible, additive only)
- **Storage drivers**: Need to implement container-type checks in delete/move handlers
- **Gateway**: Needs to pass through new permission fields
- **Clients**: Can gradually adopt new fields; existing behavior unchanged
