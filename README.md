# casbin

Self-research about Casbin with ACL Model

## Casbin Feature set for Go:

- Enforcement
- RBAC
- RBAC API
- ABAC
- Role Manager
- Adapter
- Filtered Adapter
- Management API
- Watcher
- Multi-Threading

Watcher or Role Manager only means having the interface in the core library.

## What Casbin does:

1. Enforce the policy in the classic {subject, object, action} form or a customized form as you defined, both allow and deny authorizations are supported.
2. Handle the storage of the access control model and its policy.
3. Manage the role-user mappings and role-role mappings (aka role hierarchy in RBAC).
4. Support built-in superuser like root or administrator. A superuser can do anything without explicit permissions.
5. Multiple built-in operators to support the rule matching. For example, keyMatch can map a resource key /foo/bar to the pattern /foo\*.

## What Casbin does NOT do:

1. Authentication (aka verify username and password when a user logs in)
2. Manage the list of users or roles. I believe it's more convenient for the project itself to manage these entities. Users usually have their passwords, and Casbin is not designed as a password container. However, Casbin stores the user-role mapping for the RBAC scenario
