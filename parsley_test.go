package parsley

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/scottkgregory/parsley/internal/assert"
)

func TestParse(t *testing.T) {
	data := map[string]any{}
	err := json.Unmarshal([]byte(`{
  "object_kind": "build",
  "ref": "gitlab-script-trigger",
  "tag": false,
  "before_sha": "2293ada6b400935a1378653304eaf6221e0fdb8f",
  "sha": "2293ada6b400935a1378653304eaf6221e0fdb8f",
  "build_id": 1977,
  "build_name": "test",
  "build_stage": "test",
  "build_status": "created",
  "build_created_at": "2021-02-23T02:41:37.886Z",
  "build_created_at_iso": "2021-02-23T02:41:37Z",
  "build_started_at": null,
  "build_started_at_iso": null,
  "build_finished_at": null,
  "build_finished_at_iso": null,
  "build_duration": null,
  "build_queued_duration": 1095.588715,
  "build_allow_failure": false,
  "build_failure_reason": "script_failure",
  "retries_count": 2,
  "pipeline_id": 2366,
  "project_id": 380,
  "project_name": "gitlab-org/gitlab-test",
  "user": {
    "id": 3,
    "name": "User",
    "email": "user@gitlab.com",
    "avatar_url": "http://www.gravatar.com/avatar/e32bd13e2add097461cb96824b7a829c?s=80\u0026d=identicon"
  },
  "commit": {
    "id": 2366,
    "name": "Build pipeline",
    "sha": "2293ada6b400935a1378653304eaf6221e0fdb8f",
    "message": "test\n",
    "author_name": "User",
    "author_email": "user@gitlab.com",
    "status": "created",
    "duration": null,
    "started_at": null,
    "started_at_iso": null,
    "finished_at": null,
    "finished_at_iso": null
  },
  "repository": {
    "name": "gitlab_test",
    "description": "Atque in sunt eos similique dolores voluptatem.",
    "homepage": "http://192.168.64.1:3005/gitlab-org/gitlab-test",
    "git_ssh_url": "git@192.168.64.1:gitlab-org/gitlab-test.git",
    "git_http_url": "http://192.168.64.1:3005/gitlab-org/gitlab-test.git",
    "visibility_level": 20
  },
  "project": {
    "id": 380,
    "name": "Gitlab Test",
    "description": "Atque in sunt eos similique dolores voluptatem.",
    "web_url": "http://192.168.64.1:3005/gitlab-org/gitlab-test",
    "avatar_url": null,
    "git_ssh_url": "git@192.168.64.1:gitlab-org/gitlab-test.git",
    "git_http_url": "http://192.168.64.1:3005/gitlab-org/gitlab-test.git",
    "namespace": "Gitlab Org",
    "visibility_level": 20,
    "path_with_namespace": "gitlab-org/gitlab-test",
    "default_branch": "master"
  },
  "runner": {
    "active": true,
    "runner_type": "project_type",
    "is_shared": false,
    "id": 380987,
    "description": "shared-runners-manager-6.gitlab.com",
    "tags": ["linux", "docker"]
  },
  "environment": null,
  "source_pipeline": {
    "project": {
      "id": 41,
      "web_url": "https://gitlab.example.com/gitlab-org/upstream-project",
      "path_with_namespace": "gitlab-org/upstream-project"
    },
    "pipeline_id": 30,
    "job_id": 3401
  }
}
`), &data)
	if err != nil {
		panic(err)
	}

	testCases := []struct {
		name           string
		input          string
		data           map[string]any
		expectedBool   *bool
		expectedAny    any
		expectedString any
	}{
		{
			name:           "basic negation",
			input:          "-2",
			data:           map[string]any{},
			expectedBool:   toPtr(false),
			expectedAny:    float64(-2),
			expectedString: "-2",
		},
		{
			name:           "negation",
			input:          "-(2*2)",
			data:           map[string]any{},
			expectedBool:   toPtr(false),
			expectedAny:    float64(-4),
			expectedString: "-4",
		},
		{
			name:           "basic addition",
			input:          "1+2",
			data:           map[string]any{},
			expectedBool:   toPtr(true),
			expectedAny:    float64(3),
			expectedString: "3",
		},
		{
			name:           "basic subtraction",
			input:          "10-1",
			data:           map[string]any{},
			expectedBool:   toPtr(true),
			expectedAny:    float64(9),
			expectedString: "9",
		},
		{
			name:           "basic division",
			input:          "3/3",
			data:           map[string]any{},
			expectedBool:   toPtr(true),
			expectedAny:    float64(1),
			expectedString: "1",
		},
		{
			name:           "basic division with brackets",
			input:          "(3+3)/3",
			data:           map[string]any{},
			expectedBool:   toPtr(true),
			expectedAny:    float64(2),
			expectedString: "2",
		},
		{
			name:           "basic division with brackets",
			input:          "(3/3)+5",
			data:           map[string]any{},
			expectedBool:   toPtr(true),
			expectedAny:    float64(6),
			expectedString: "6",
		},
		{
			name:           "const equality int",
			input:          "5 == 5",
			data:           map[string]any{},
			expectedBool:   toPtr(true),
			expectedAny:    true,
			expectedString: "true",
		},
		{
			name:           "variable equality int",
			input:          "foo == 5",
			data:           map[string]any{"foo": 5},
			expectedBool:   toPtr(true),
			expectedAny:    true,
			expectedString: "true",
		},
		{
			name:           "variable equality int (not equal)",
			input:          "foo == 5",
			data:           map[string]any{"foo": 6},
			expectedBool:   toPtr(false),
			expectedAny:    false,
			expectedString: "false",
		},
		{
			name:           "variable equality string",
			input:          `foo == "hello"`,
			data:           map[string]any{"foo": "hello"},
			expectedBool:   toPtr(true),
			expectedAny:    true,
			expectedString: "true",
		},
		{
			name:           "deep variable equality string",
			input:          "foo.bar.baz == \"hello\"",
			data:           map[string]any{"foo": map[string]any{"bar": map[string]any{"baz": "hello"}}},
			expectedBool:   toPtr(true),
			expectedAny:    true,
			expectedString: "true",
		},
		{
			name:           "deep variable equality int",
			input:          "foo.bar.baz == 6",
			data:           map[string]any{"foo": map[string]any{"bar": map[string]any{"baz": 6}}},
			expectedBool:   toPtr(true),
			expectedAny:    true,
			expectedString: "true",
		},
		{
			name:           "math ceil",
			input:          `ceil(2.1)`,
			data:           map[string]any{},
			expectedBool:   toPtr(true),
			expectedAny:    float64(3),
			expectedString: "3",
		},
		{
			name:           "math floor",
			input:          `floor(2.9)`,
			data:           map[string]any{},
			expectedBool:   toPtr(true),
			expectedAny:    float64(2),
			expectedString: "2",
		},
		{
			name:           "math round up",
			input:          `round(2.9)`,
			data:           map[string]any{},
			expectedBool:   toPtr(true),
			expectedAny:    float64(3),
			expectedString: "3",
		},
		{
			name:           "math round down",
			input:          `round(2.49)`,
			data:           map[string]any{},
			expectedBool:   toPtr(true),
			expectedAny:    float64(2),
			expectedString: "2",
		},
		{
			name:           "math truncate",
			input:          `truncate(2.9)`,
			data:           map[string]any{},
			expectedBool:   toPtr(true),
			expectedAny:    float64(2),
			expectedString: "2",
		},
		{
			name:           "math absolute",
			input:          `absolute(2.9)`,
			data:           map[string]any{},
			expectedBool:   toPtr(true),
			expectedAny:    2.9,
			expectedString: "2.9",
		},
		{
			name:           "greater than",
			input:          `1 > 2`,
			data:           map[string]any{},
			expectedBool:   toPtr(false),
			expectedAny:    false,
			expectedString: "false",
		},
		{
			name:           "greater than",
			input:          `2 > 1`,
			data:           map[string]any{},
			expectedBool:   toPtr(true),
			expectedAny:    true,
			expectedString: "true",
		},
		{
			name:           "less than",
			input:          `2 < 1`,
			data:           map[string]any{},
			expectedBool:   toPtr(false),
			expectedAny:    false,
			expectedString: "false",
		},
		{
			name:           "less than",
			input:          `1 < 2`,
			data:           map[string]any{},
			expectedBool:   toPtr(true),
			expectedAny:    true,
			expectedString: "true",
		},
		{
			name:           "less than",
			input:          `foo < 2`,
			data:           map[string]any{"foo": 1.5},
			expectedBool:   toPtr(true),
			expectedAny:    true,
			expectedString: "true",
		},
		{
			name:           "less than and greater than",
			input:          `(foo < 2) && (foo > 1)`,
			data:           map[string]any{"foo": 1.5},
			expectedBool:   toPtr(true),
			expectedAny:    true,
			expectedString: "true",
		},
		{
			name:           "less than and greater than",
			input:          `(foo < 2) && (foo > 1)`,
			data:           map[string]any{"foo": 4},
			expectedBool:   toPtr(false),
			expectedAny:    false,
			expectedString: "false",
		},
		{
			name:           "less than or greater than",
			input:          `(foo < 2) || (foo > 1)`,
			data:           map[string]any{"foo": 1.5},
			expectedBool:   toPtr(true),
			expectedAny:    true,
			expectedString: "true",
		},
		{
			name:           "less than or greater than",
			input:          `(foo < 2) || (foo > 1)`,
			data:           map[string]any{"foo": 4},
			expectedBool:   toPtr(true),
			expectedAny:    true,
			expectedString: "true",
		},
		{
			name:           "gitlab sample match",
			input:          `(object_attributes.state == "opened") && (object_attributes.labels.title == "automated")`,
			data:           map[string]any{"object_attributes": map[string]any{"state": "opened", "labels": map[string]any{"title": "automated"}}},
			expectedBool:   toPtr(true),
			expectedAny:    true,
			expectedString: "true",
		},
		{
			name:           "gitlab sample no match",
			input:          `(object_attributes.state == "opened") && (object_attributes.labels.title == "automated")`,
			data:           map[string]any{"object_attributes": map[string]any{"state": "opened", "labels": map[string]any{"title": "made by gary"}}},
			expectedBool:   toPtr(false),
			expectedAny:    false,
			expectedString: "false",
		},
		{
			name:           "contains_any (true)",
			input:          `(object_attributes.state == "opened") && contains_any(object_attributes.labels, "title", "automated")`,
			data:           map[string]any{"object_attributes": map[string]any{"state": "opened", "labels": []any{map[string]any{"title": "automated"}}}},
			expectedBool:   toPtr(true),
			expectedAny:    true,
			expectedString: "true",
		},
		{
			name:           "contains_any (false)",
			input:          `(object_attributes.state == "opened") && contains_any(object_attributes.labels, "title", "automated")`,
			data:           map[string]any{"object_attributes": map[string]any{"state": "opened", "labels": []any{map[string]any{"title": "made by gary"}}}},
			expectedBool:   toPtr(false),
			expectedAny:    false,
			expectedString: "false",
		},
		{
			name:           "contains_any number (true)",
			input:          `(object_attributes.state == "opened") && contains_any(object_attributes.labels, "title", 70)`,
			data:           map[string]any{"object_attributes": map[string]any{"state": "opened", "labels": []any{map[string]any{"title": 70}}}},
			expectedBool:   toPtr(true),
			expectedAny:    true,
			expectedString: "true",
		},
		{
			name:           "contains_any number (false)",
			input:          `(object_attributes.state == "opened") && contains_any(object_attributes.labels, "title", 70)`,
			data:           map[string]any{"object_attributes": map[string]any{"state": "opened", "labels": []any{map[string]any{"title": 80}}}},
			expectedBool:   toPtr(false),
			expectedAny:    false,
			expectedString: "false",
		},
		{
			name:           "prints data",
			input:          `object_attributes.state`,
			data:           map[string]any{"object_attributes": map[string]any{"state": "opened", "labels": []any{map[string]any{"title": 80}}}},
			expectedBool:   nil,
			expectedAny:    "opened",
			expectedString: "opened",
		},
		{
			name:           "equal, no data",
			input:          `event_type == "merge_request"`,
			data:           map[string]any{},
			expectedBool:   toPtr(false),
			expectedAny:    false,
			expectedString: "false",
		},
		{
			name:           "equal, no data",
			input:          `(event_type == "merge_request") && contains_any(object_attributes.labels, "title", "automated")`,
			data:           data,
			expectedBool:   toPtr(false),
			expectedAny:    false,
			expectedString: "false",
		},
		{
			name:  "and not",
			input: `(event_type == "note") && (not(contains_any(merge_request.labels, "title", "automated")))`,
			data: map[string]any{
				"event_type": "note",
				"merge_request": map[string]any{
					"labels": []any{map[string]any{"title": "automated"}},
				},
			},
			expectedBool:   toPtr(false),
			expectedAny:    false,
			expectedString: "false",
		},
		{
			name:  "and not where data does not exist",
			input: `(event_type == "note") && (not(contains_any(merge_request.labels, "title", "automated")))`,
			data: map[string]any{
				"event_type": "note",
			},
			expectedBool:   toPtr(true),
			expectedAny:    true,
			expectedString: "true",
		},
		{
			name:  "concat",
			input: `event_type + foo`,
			data: map[string]any{
				"event_type": "note",
				"foo":        "_",
			},
			expectedBool:   nil,
			expectedAny:    "note_",
			expectedString: "note_",
		},
	}

	for _, tc := range testCases {
		if tc.expectedBool != nil {
			t.Run(fmt.Sprintf("as_bool_%s", tc.name), func(tt *testing.T) {
				matcher, err := NewParser(true)
				assert.Nil(t, err)

				actual, err := matcher.ParseAsBool(tc.input, tc.data)
				assert.Equal(tt, *tc.expectedBool, actual)
				assert.Nil(t, err)

				matcher.Close()
			})
		}

		t.Run(fmt.Sprintf("as_string_%s", tc.name), func(tt *testing.T) {
			matcher, err := NewParser(false)
			assert.Nil(t, err)

			actual, err := matcher.ParseAsString(tc.input, tc.data)
			assert.Equal(tt, tc.expectedString, actual)
			assert.Nil(t, err)

			matcher.Close()
		})

		t.Run(fmt.Sprintf("as_any_%s", tc.name), func(tt *testing.T) {
			matcher, err := NewParser(false)
			assert.Nil(t, err)

			actual, err := matcher.ParseAsAny(tc.input, tc.data)
			assert.Equal(tt, tc.expectedAny, actual)
			assert.Nil(t, err)

			matcher.Close()
		})
	}
}

func toPtr[T any](t T) *T {
	return &t
}
