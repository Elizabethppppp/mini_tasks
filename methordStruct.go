package main

func Clone(g Group) Group {
	name := g.Name

	members := make([]string, len(g.Members))
	copy(members, g.Members)

	return Group{
		Name: name,
		Members: members,
	}
}
