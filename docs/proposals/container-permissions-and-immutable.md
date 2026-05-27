# Proposal: Container-Specific Permissions and Immutable Resources

## Status

Draft

## Authors

- @flash7777

## Summary

This proposal adds:

1. **`delete_container`** + **`move_container`** — two new permission fields in `ResourcePermissions` to distinguish file and container operations
2. **`immutable`** — a persistent attribute on `ResourceInfo` plus `SetImmutable`/`UnsetImmutable` RPCs to freeze resources

## Motivation

### Problem 1: No distinction between file and container operations

The current `ResourcePermissions` uses a single `delete` flag for both files and containers, and a single `move` flag for both. This makes it impossible to configure a role that allows file deletion but protects directory structures.

**Real-world example: File plans (Aktenplan)**

In German public administration, a rigid hierarchical directory structure (Aktenplan) must be maintained. Users work with documents freely, but must not alter the directory structure.

```
/01 Administration          <- protected (no delete/rename)
  /01.01 HR                 <- protected
    /01.01.01 Recruiting    <- protected
      /Applications/        <- user can create, edit, delete files here
        resume.pdf
```

### Problem 2: No immutability concept

Locks are temporary (expiration-based) and designed for collaborative editing. What is missing is a **persistent attribute** on a resource — stored as xattr, returned via `Stat()`, set/unset via dedicated RPCs — that prevents structural changes.

This is fundamentally different from a transient processing status (cf. #190). Immutable is a deliberate administrative decision, not a system state.

## Specification

### New fields in ResourcePermissions

```protobuf
message ResourcePermissions {
  // ... existing fields 1-20 ...

  bool delete_container = 21;
  bool move_container = 22;
}
```

- `delete_container`: controls whether containers can be deleted, independent of `delete` (which then applies to files only)
- `move_container`: controls whether containers can be moved/renamed, independent of `move`

### Backward compatibility

When not explicitly set, implementations SHOULD fall back to `delete`/`move` for containers. Existing clients work unchanged.

### Immutable attribute on ResourceInfo

```protobuf
message ResourceInfo {
  // ... existing fields 1-19 ...

  bool immutable = 20;
}
```

This is a **persistent attribute** (stored as xattr on the filesystem), not a transient status. It is returned by `Stat()` and set/cleared via dedicated RPCs.

Semantics:
- **Files**: cannot be modified, deleted, moved or renamed
- **Containers**: additionally prevents creation of new children; existing non-immutable children can still be modified

### RPCs

```protobuf
rpc SetImmutable(SetImmutableRequest) returns (SetImmutableResponse);
rpc UnsetImmutable(UnsetImmutableRequest) returns (UnsetImmutableResponse);
```

Only users with management permissions (space owner, space manager, administrator) may call these RPCs.

The pattern follows `SetLock`/`Unlock` — a `Reference` identifies the target resource.

### Distinction: attribute vs. action

| | Attribute (`immutable`) | Action (`SetImmutable` / `UnsetImmutable`) |
|---|---|---|
| What | Boolean on the resource, persisted as xattr | RPC to change the attribute |
| When read | Returned in `ResourceInfo` via `Stat()` | — |
| Who sets | Manager / Admin via RPC | — |
| Expiration | None (permanent until explicitly unset) | — |

### Distinction from Locks (#190 Status)

| | Lock | Immutable | Status (#190) |
|---|---|---|---|
| Purpose | Collaborative editing | Structure/content protection | Processing state |
| Duration | Temporary (with expiration) | Permanent (until admin unsets) | Transient |
| Storage | Lock file / xattr | xattr | Opaque / xattr |
| Set by | Any user with write access | Manager / Admin only | System / App |
| Scope | Prevents concurrent writes | Prevents all modifications | Informational |

## ACL representation

```
+dc  = delete_container allowed
!dc  = delete_container denied
+mc  = move_container allowed
!mc  = move_container denied
```

## Use cases

1. **File plan (Aktenplan) protection**: rigid directory structure, free file operations below
2. **Retention / Legal hold**: freeze resources during retention periods
3. **Archive directories**: freeze completed project folders
4. **Collaborative workspaces**: prevent accidental directory deletion
5. **Compliance**: regulatory requirements for records management

## Impact

- **Proto changes**: 2 new fields in ResourcePermissions, 1 new field in ResourceInfo, 2 new RPCs (all additive, backward compatible)
- **Storage drivers**: container-type checks in delete/move handlers; xattr for immutable
- **Gateway**: pass through new fields and RPCs
- **Clients**: gradual adoption; existing behavior unchanged
