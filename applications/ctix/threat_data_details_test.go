package ctix

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestThreatDataRelationsResponseAcceptsTargetTimestamps(t *testing.T) {
	tests := []struct {
		name     string
		created  string
		modified string
	}{
		{name: "integer timestamps", created: "1730000000", modified: "1730000001"},
		{name: "fractional created timestamp", created: "1730000000.125", modified: "1730000001"},
		{name: "fractional modified timestamp", created: "1730000000", modified: "1730000001.25"},
		{name: "fractional timestamps", created: "1730000000.125", modified: "1730000001.25"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payload := `{
				"next": "next-page",
				"page_size": 1,
				"previous": "previous-page",
				"results": [{
					"created_by_ref": null,
					"id": "relation-id",
					"is_forward": true,
					"is_redacted": false,
					"relation_end": null,
					"relation_start": null,
					"relationship_confidence": null,
					"relationship_type": "related-to",
					"sources": [{"id": "source-id", "name": "source", "source_type": "feed"}],
					"target_ref": {
						"created": ` + tt.created + `,
						"id": "target-id",
						"modified": ` + tt.modified + `,
						"name": "target-name",
						"object_type": "indicator",
						"sub_type": null,
						"tlp": "CLEAR"
					}
				}],
				"total": 1
			}`

			var response ThreatDataRelationsResponse
			if err := json.Unmarshal([]byte(payload), &response); err != nil {
				t.Fatalf("decode relation response: %v", err)
			}

			if response.Next != "next-page" || response.Previous != "previous-page" {
				t.Fatalf("pagination links were not preserved: next=%v previous=%v", response.Next, response.Previous)
			}
			if response.PageSize != 1 || response.Total != 1 || len(response.Results) != 1 {
				t.Fatalf("page metadata or relation count changed: page_size=%d total=%d results=%d", response.PageSize, response.Total, len(response.Results))
			}

			relation := response.Results[0]
			if relation.ID != "relation-id" || relation.RelationshipType != "related-to" || relation.TargetRef.ID != "target-id" {
				t.Fatalf("relation fields were not preserved: %+v", relation)
			}
			if relation.TargetRef.Created.String() != tt.created {
				t.Fatalf("target created timestamp = %q, want %q", relation.TargetRef.Created, tt.created)
			}
			if relation.TargetRef.Modified.String() != tt.modified {
				t.Fatalf("target modified timestamp = %q, want %q", relation.TargetRef.Modified, tt.modified)
			}

			encoded, err := json.Marshal(response)
			if err != nil {
				t.Fatalf("encode relation response: %v", err)
			}
			for _, timestamp := range []struct {
				field string
				value string
			}{
				{field: "created", value: tt.created},
				{field: "modified", value: tt.modified},
			} {
				if !strings.Contains(string(encoded), `"`+timestamp.field+`":`+timestamp.value) {
					t.Fatalf("encoded response did not preserve numeric %s timestamp %q: %s", timestamp.field, timestamp.value, encoded)
				}
			}
		})
	}
}
