package domain

type CategoryKind string

const (
	MAIN_CATEGORY CategoryKind = "Main Category"
	CATEGORY CategoryKind = "Category"
	LEAF_CATEGORY CategoryKind = "Leaf Category"
)