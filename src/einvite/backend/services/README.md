backend.services
==
This package contains the services invoked by the frontend layer for database access or integration with external services.

The decision of having services and repositories is that repositories only contains the database access logic, while services use the repositories and may need to invoke external services for authentication, get provider based content and such.