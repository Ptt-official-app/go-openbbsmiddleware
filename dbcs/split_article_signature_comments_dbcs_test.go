package dbcs

import (
	"reflect"
	"testing"
)

func Test_splitArticleSignatureCommentsDBCS(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		content []byte
	}
	tests := []struct {
		name                  string
		args                  args
		expectedArticleDBCS   []byte
		expectedSignatureDBCS []byte
		expectedComments      []byte
	}{
		// TODO: Add test cases.
		{
			name:                  "0_" + testFilename0,
			args:                  args{content: testContentAll0},
			expectedArticleDBCS:   testContent0,
			expectedSignatureDBCS: testSignature0,
			expectedComments:      testComment0,
		},
		{
			name:                  "1_" + testFilename1,
			args:                  args{content: testContentAll1},
			expectedArticleDBCS:   testContent1,
			expectedSignatureDBCS: testSignature1,
			expectedComments:      testComment1,
		},
		{
			name:                  "2_" + testFilename2,
			args:                  args{content: testContentAll2},
			expectedArticleDBCS:   testContent2,
			expectedSignatureDBCS: testSignature2,
			expectedComments:      testComment2,
		},
		{
			name:                  "3_" + testFilename3,
			args:                  args{content: testContentAll3},
			expectedArticleDBCS:   testContent3,
			expectedSignatureDBCS: testSignature3,
			expectedComments:      testComment3,
		},
		{
			name:                  "4_" + testFilename4,
			args:                  args{content: testContentAll4},
			expectedArticleDBCS:   testContent4,
			expectedSignatureDBCS: testSignature4,
			expectedComments:      testComment4,
		},
		{
			name:                  "5_" + testFilename5,
			args:                  args{content: testContentAll5},
			expectedArticleDBCS:   testContent5,
			expectedSignatureDBCS: testSignature5,
			expectedComments:      testComment5,
		},
		{
			name:                  "6_" + testFilename6,
			args:                  args{content: testContentAll6},
			expectedArticleDBCS:   testContent6,
			expectedSignatureDBCS: testSignature6,
			expectedComments:      testComment6,
		},
		{
			name:                  "7_" + testFilename7,
			args:                  args{content: testContentAll7},
			expectedArticleDBCS:   testContent7,
			expectedSignatureDBCS: testSignature7,
			expectedComments:      testComment7,
		},
		{
			name:                  "8_" + testFilename8,
			args:                  args{content: testContentAll8},
			expectedArticleDBCS:   testContent8,
			expectedSignatureDBCS: testSignature8,
			expectedComments:      testComment8,
		},
		{
			name:                  "9_" + testFilename9,
			args:                  args{content: testContentAll9},
			expectedArticleDBCS:   testContent9,
			expectedSignatureDBCS: testSignature9,
			expectedComments:      testComment9,
		},
		{
			name:                  "10_" + testFilename10,
			args:                  args{content: testContentAll10},
			expectedArticleDBCS:   testContent10,
			expectedSignatureDBCS: testSignature10,
			expectedComments:      testComment10,
		},
		{
			name:                  "11_" + testFilename11,
			args:                  args{content: testContentAll11},
			expectedArticleDBCS:   testContent11,
			expectedSignatureDBCS: testSignature11,
			expectedComments:      testComment11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotArticleDBCS, gotSignatureDBCS, gotComments := splitArticleSignatureCommentsDBCS(tt.args.content)
			if !reflect.DeepEqual(gotArticleDBCS, tt.expectedArticleDBCS) {
				t.Errorf("splitArticleSignatureCommentsDBCS() gotArticleDBCS = %v, want %v", gotArticleDBCS, tt.expectedArticleDBCS)
			}
			if !reflect.DeepEqual(gotSignatureDBCS, tt.expectedSignatureDBCS) {
				t.Errorf("splitArticleSignatureCommentsDBCS() gotSignatureDBCS = %v, want %v", gotSignatureDBCS, tt.expectedSignatureDBCS)
			}
			if !reflect.DeepEqual(gotComments, tt.expectedComments) {
				t.Errorf("splitArticleSignatureCommentsDBCS() gotComments = %v, want %v", gotComments, tt.expectedComments)
			}
		})
	}
}
