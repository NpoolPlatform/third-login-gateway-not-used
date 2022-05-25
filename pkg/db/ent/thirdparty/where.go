// Code generated by entc, DO NOT EDIT.

package thirdparty

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/third-login-gateway/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// BrandName applies equality check predicate on the "brand_name" field. It's identical to BrandNameEQ.
func BrandName(v string) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBrandName), v))
	})
}

// Logo applies equality check predicate on the "logo" field. It's identical to LogoEQ.
func Logo(v string) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLogo), v))
	})
}

// Domain applies equality check predicate on the "domain" field. It's identical to DomainEQ.
func Domain(v string) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDomain), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.ThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ThirdParty(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.ThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ThirdParty(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.ThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ThirdParty(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.ThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ThirdParty(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.ThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ThirdParty(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.ThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ThirdParty(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// BrandNameEQ applies the EQ predicate on the "brand_name" field.
func BrandNameEQ(v string) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBrandName), v))
	})
}

// BrandNameNEQ applies the NEQ predicate on the "brand_name" field.
func BrandNameNEQ(v string) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBrandName), v))
	})
}

// BrandNameIn applies the In predicate on the "brand_name" field.
func BrandNameIn(vs ...string) predicate.ThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ThirdParty(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBrandName), v...))
	})
}

// BrandNameNotIn applies the NotIn predicate on the "brand_name" field.
func BrandNameNotIn(vs ...string) predicate.ThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ThirdParty(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBrandName), v...))
	})
}

// BrandNameGT applies the GT predicate on the "brand_name" field.
func BrandNameGT(v string) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBrandName), v))
	})
}

// BrandNameGTE applies the GTE predicate on the "brand_name" field.
func BrandNameGTE(v string) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBrandName), v))
	})
}

// BrandNameLT applies the LT predicate on the "brand_name" field.
func BrandNameLT(v string) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBrandName), v))
	})
}

// BrandNameLTE applies the LTE predicate on the "brand_name" field.
func BrandNameLTE(v string) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBrandName), v))
	})
}

// BrandNameContains applies the Contains predicate on the "brand_name" field.
func BrandNameContains(v string) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBrandName), v))
	})
}

// BrandNameHasPrefix applies the HasPrefix predicate on the "brand_name" field.
func BrandNameHasPrefix(v string) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBrandName), v))
	})
}

// BrandNameHasSuffix applies the HasSuffix predicate on the "brand_name" field.
func BrandNameHasSuffix(v string) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBrandName), v))
	})
}

// BrandNameEqualFold applies the EqualFold predicate on the "brand_name" field.
func BrandNameEqualFold(v string) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBrandName), v))
	})
}

// BrandNameContainsFold applies the ContainsFold predicate on the "brand_name" field.
func BrandNameContainsFold(v string) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBrandName), v))
	})
}

// LogoEQ applies the EQ predicate on the "logo" field.
func LogoEQ(v string) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLogo), v))
	})
}

// LogoNEQ applies the NEQ predicate on the "logo" field.
func LogoNEQ(v string) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLogo), v))
	})
}

// LogoIn applies the In predicate on the "logo" field.
func LogoIn(vs ...string) predicate.ThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ThirdParty(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLogo), v...))
	})
}

// LogoNotIn applies the NotIn predicate on the "logo" field.
func LogoNotIn(vs ...string) predicate.ThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ThirdParty(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLogo), v...))
	})
}

// LogoGT applies the GT predicate on the "logo" field.
func LogoGT(v string) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLogo), v))
	})
}

// LogoGTE applies the GTE predicate on the "logo" field.
func LogoGTE(v string) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLogo), v))
	})
}

// LogoLT applies the LT predicate on the "logo" field.
func LogoLT(v string) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLogo), v))
	})
}

// LogoLTE applies the LTE predicate on the "logo" field.
func LogoLTE(v string) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLogo), v))
	})
}

// LogoContains applies the Contains predicate on the "logo" field.
func LogoContains(v string) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLogo), v))
	})
}

// LogoHasPrefix applies the HasPrefix predicate on the "logo" field.
func LogoHasPrefix(v string) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLogo), v))
	})
}

// LogoHasSuffix applies the HasSuffix predicate on the "logo" field.
func LogoHasSuffix(v string) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLogo), v))
	})
}

// LogoEqualFold applies the EqualFold predicate on the "logo" field.
func LogoEqualFold(v string) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLogo), v))
	})
}

// LogoContainsFold applies the ContainsFold predicate on the "logo" field.
func LogoContainsFold(v string) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLogo), v))
	})
}

// DomainEQ applies the EQ predicate on the "domain" field.
func DomainEQ(v string) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDomain), v))
	})
}

// DomainNEQ applies the NEQ predicate on the "domain" field.
func DomainNEQ(v string) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDomain), v))
	})
}

// DomainIn applies the In predicate on the "domain" field.
func DomainIn(vs ...string) predicate.ThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ThirdParty(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDomain), v...))
	})
}

// DomainNotIn applies the NotIn predicate on the "domain" field.
func DomainNotIn(vs ...string) predicate.ThirdParty {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ThirdParty(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDomain), v...))
	})
}

// DomainGT applies the GT predicate on the "domain" field.
func DomainGT(v string) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDomain), v))
	})
}

// DomainGTE applies the GTE predicate on the "domain" field.
func DomainGTE(v string) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDomain), v))
	})
}

// DomainLT applies the LT predicate on the "domain" field.
func DomainLT(v string) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDomain), v))
	})
}

// DomainLTE applies the LTE predicate on the "domain" field.
func DomainLTE(v string) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDomain), v))
	})
}

// DomainContains applies the Contains predicate on the "domain" field.
func DomainContains(v string) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDomain), v))
	})
}

// DomainHasPrefix applies the HasPrefix predicate on the "domain" field.
func DomainHasPrefix(v string) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDomain), v))
	})
}

// DomainHasSuffix applies the HasSuffix predicate on the "domain" field.
func DomainHasSuffix(v string) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDomain), v))
	})
}

// DomainEqualFold applies the EqualFold predicate on the "domain" field.
func DomainEqualFold(v string) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDomain), v))
	})
}

// DomainContainsFold applies the ContainsFold predicate on the "domain" field.
func DomainContainsFold(v string) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDomain), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ThirdParty) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ThirdParty) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ThirdParty) predicate.ThirdParty {
	return predicate.ThirdParty(func(s *sql.Selector) {
		p(s.Not())
	})
}
