// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package authz

import (
	"context"

	"github.com/facebookincubator/symphony/graph/ent/hyperlink"

	"github.com/facebookincubator/symphony/graph/ent"
	"github.com/facebookincubator/symphony/graph/ent/privacy"
)

// HyperlinkReadPolicyRule grants read permission to hyperlink based on policy.
func HyperlinkReadPolicyRule() privacy.QueryRule {
	return privacy.HyperlinkQueryRuleFunc(func(ctx context.Context, q *ent.HyperlinkQuery) error {
		woPredicate := workOrderReadPredicate(ctx)
		if woPredicate != nil {
			q.Where(
				hyperlink.Or(
					hyperlink.Not(hyperlink.HasWorkOrder()),
					hyperlink.HasWorkOrderWith(woPredicate)))
		}
		return privacy.Skip
	})
}
