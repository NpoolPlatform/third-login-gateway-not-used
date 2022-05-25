// Code generated by entc, DO NOT EDIT.

package auth

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/third-login-gateway/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
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
func IDNotIn(ids ...uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
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
func IDGT(id uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// ThirdPartyID applies equality check predicate on the "third_party_id" field. It's identical to ThirdPartyIDEQ.
func ThirdPartyID(v uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThirdPartyID), v))
	})
}

// AppKey applies equality check predicate on the "app_key" field. It's identical to AppKeyEQ.
func AppKey(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppKey), v))
	})
}

// AppSecret applies equality check predicate on the "app_secret" field. It's identical to AppSecretEQ.
func AppSecret(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppSecret), v))
	})
}

// RedirectURL applies equality check predicate on the "redirect_url" field. It's identical to RedirectURLEQ.
func RedirectURL(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRedirectURL), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.Auth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Auth(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...uint32) predicate.Auth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Auth(func(s *sql.Selector) {
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
func CreatedAtGT(v uint32) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.Auth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Auth(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...uint32) predicate.Auth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Auth(func(s *sql.Selector) {
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
func UpdatedAtGT(v uint32) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.Auth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Auth(func(s *sql.Selector) {
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
func DeletedAtNotIn(vs ...uint32) predicate.Auth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Auth(func(s *sql.Selector) {
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
func DeletedAtGT(v uint32) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppID), v))
	})
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...uuid.UUID) predicate.Auth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Auth(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAppID), v...))
	})
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...uuid.UUID) predicate.Auth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Auth(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAppID), v...))
	})
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppID), v))
	})
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppID), v))
	})
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppID), v))
	})
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppID), v))
	})
}

// ThirdPartyIDEQ applies the EQ predicate on the "third_party_id" field.
func ThirdPartyIDEQ(v uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThirdPartyID), v))
	})
}

// ThirdPartyIDNEQ applies the NEQ predicate on the "third_party_id" field.
func ThirdPartyIDNEQ(v uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldThirdPartyID), v))
	})
}

// ThirdPartyIDIn applies the In predicate on the "third_party_id" field.
func ThirdPartyIDIn(vs ...uuid.UUID) predicate.Auth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Auth(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldThirdPartyID), v...))
	})
}

// ThirdPartyIDNotIn applies the NotIn predicate on the "third_party_id" field.
func ThirdPartyIDNotIn(vs ...uuid.UUID) predicate.Auth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Auth(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldThirdPartyID), v...))
	})
}

// ThirdPartyIDGT applies the GT predicate on the "third_party_id" field.
func ThirdPartyIDGT(v uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldThirdPartyID), v))
	})
}

// ThirdPartyIDGTE applies the GTE predicate on the "third_party_id" field.
func ThirdPartyIDGTE(v uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldThirdPartyID), v))
	})
}

// ThirdPartyIDLT applies the LT predicate on the "third_party_id" field.
func ThirdPartyIDLT(v uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldThirdPartyID), v))
	})
}

// ThirdPartyIDLTE applies the LTE predicate on the "third_party_id" field.
func ThirdPartyIDLTE(v uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldThirdPartyID), v))
	})
}

// AppKeyEQ applies the EQ predicate on the "app_key" field.
func AppKeyEQ(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppKey), v))
	})
}

// AppKeyNEQ applies the NEQ predicate on the "app_key" field.
func AppKeyNEQ(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppKey), v))
	})
}

// AppKeyIn applies the In predicate on the "app_key" field.
func AppKeyIn(vs ...string) predicate.Auth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Auth(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAppKey), v...))
	})
}

// AppKeyNotIn applies the NotIn predicate on the "app_key" field.
func AppKeyNotIn(vs ...string) predicate.Auth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Auth(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAppKey), v...))
	})
}

// AppKeyGT applies the GT predicate on the "app_key" field.
func AppKeyGT(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppKey), v))
	})
}

// AppKeyGTE applies the GTE predicate on the "app_key" field.
func AppKeyGTE(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppKey), v))
	})
}

// AppKeyLT applies the LT predicate on the "app_key" field.
func AppKeyLT(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppKey), v))
	})
}

// AppKeyLTE applies the LTE predicate on the "app_key" field.
func AppKeyLTE(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppKey), v))
	})
}

// AppKeyContains applies the Contains predicate on the "app_key" field.
func AppKeyContains(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAppKey), v))
	})
}

// AppKeyHasPrefix applies the HasPrefix predicate on the "app_key" field.
func AppKeyHasPrefix(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAppKey), v))
	})
}

// AppKeyHasSuffix applies the HasSuffix predicate on the "app_key" field.
func AppKeyHasSuffix(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAppKey), v))
	})
}

// AppKeyEqualFold applies the EqualFold predicate on the "app_key" field.
func AppKeyEqualFold(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAppKey), v))
	})
}

// AppKeyContainsFold applies the ContainsFold predicate on the "app_key" field.
func AppKeyContainsFold(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAppKey), v))
	})
}

// AppSecretEQ applies the EQ predicate on the "app_secret" field.
func AppSecretEQ(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppSecret), v))
	})
}

// AppSecretNEQ applies the NEQ predicate on the "app_secret" field.
func AppSecretNEQ(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppSecret), v))
	})
}

// AppSecretIn applies the In predicate on the "app_secret" field.
func AppSecretIn(vs ...string) predicate.Auth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Auth(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAppSecret), v...))
	})
}

// AppSecretNotIn applies the NotIn predicate on the "app_secret" field.
func AppSecretNotIn(vs ...string) predicate.Auth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Auth(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAppSecret), v...))
	})
}

// AppSecretGT applies the GT predicate on the "app_secret" field.
func AppSecretGT(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppSecret), v))
	})
}

// AppSecretGTE applies the GTE predicate on the "app_secret" field.
func AppSecretGTE(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppSecret), v))
	})
}

// AppSecretLT applies the LT predicate on the "app_secret" field.
func AppSecretLT(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppSecret), v))
	})
}

// AppSecretLTE applies the LTE predicate on the "app_secret" field.
func AppSecretLTE(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppSecret), v))
	})
}

// AppSecretContains applies the Contains predicate on the "app_secret" field.
func AppSecretContains(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAppSecret), v))
	})
}

// AppSecretHasPrefix applies the HasPrefix predicate on the "app_secret" field.
func AppSecretHasPrefix(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAppSecret), v))
	})
}

// AppSecretHasSuffix applies the HasSuffix predicate on the "app_secret" field.
func AppSecretHasSuffix(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAppSecret), v))
	})
}

// AppSecretEqualFold applies the EqualFold predicate on the "app_secret" field.
func AppSecretEqualFold(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAppSecret), v))
	})
}

// AppSecretContainsFold applies the ContainsFold predicate on the "app_secret" field.
func AppSecretContainsFold(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAppSecret), v))
	})
}

// RedirectURLEQ applies the EQ predicate on the "redirect_url" field.
func RedirectURLEQ(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRedirectURL), v))
	})
}

// RedirectURLNEQ applies the NEQ predicate on the "redirect_url" field.
func RedirectURLNEQ(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRedirectURL), v))
	})
}

// RedirectURLIn applies the In predicate on the "redirect_url" field.
func RedirectURLIn(vs ...string) predicate.Auth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Auth(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRedirectURL), v...))
	})
}

// RedirectURLNotIn applies the NotIn predicate on the "redirect_url" field.
func RedirectURLNotIn(vs ...string) predicate.Auth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Auth(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRedirectURL), v...))
	})
}

// RedirectURLGT applies the GT predicate on the "redirect_url" field.
func RedirectURLGT(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRedirectURL), v))
	})
}

// RedirectURLGTE applies the GTE predicate on the "redirect_url" field.
func RedirectURLGTE(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRedirectURL), v))
	})
}

// RedirectURLLT applies the LT predicate on the "redirect_url" field.
func RedirectURLLT(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRedirectURL), v))
	})
}

// RedirectURLLTE applies the LTE predicate on the "redirect_url" field.
func RedirectURLLTE(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRedirectURL), v))
	})
}

// RedirectURLContains applies the Contains predicate on the "redirect_url" field.
func RedirectURLContains(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRedirectURL), v))
	})
}

// RedirectURLHasPrefix applies the HasPrefix predicate on the "redirect_url" field.
func RedirectURLHasPrefix(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRedirectURL), v))
	})
}

// RedirectURLHasSuffix applies the HasSuffix predicate on the "redirect_url" field.
func RedirectURLHasSuffix(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRedirectURL), v))
	})
}

// RedirectURLEqualFold applies the EqualFold predicate on the "redirect_url" field.
func RedirectURLEqualFold(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRedirectURL), v))
	})
}

// RedirectURLContainsFold applies the ContainsFold predicate on the "redirect_url" field.
func RedirectURLContainsFold(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRedirectURL), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Auth) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Auth) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
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
func Not(p predicate.Auth) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		p(s.Not())
	})
}
