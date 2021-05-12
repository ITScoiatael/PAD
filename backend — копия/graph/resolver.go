package graph

import "graphql-server/prisma/db"

type Resolver struct {
	Prisma *db.PrismaClient
}
