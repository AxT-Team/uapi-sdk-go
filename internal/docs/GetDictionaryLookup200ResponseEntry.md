# GetDictionaryLookup200ResponseEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Definitions** | Pointer to [**[]GetDictionaryLookup200ResponseEntryDefinitionsInner**](GetDictionaryLookup200ResponseEntryDefinitionsInner.md) |  | [optional] 
**EnglishDefinitions** | Pointer to [**[]GetDictionaryLookup200ResponseEntryEnglishDefinitionsInner**](GetDictionaryLookup200ResponseEntryEnglishDefinitionsInner.md) |  | [optional] 
**ExamTags** | Pointer to **[]string** |  | [optional] 
**Examples** | Pointer to [**[]GetDictionaryLookup200ResponseEntryExamplesInner**](GetDictionaryLookup200ResponseEntryExamplesInner.md) |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**Phonetics** | Pointer to [**GetDictionaryLookup200ResponseEntryPhonetics**](GetDictionaryLookup200ResponseEntryPhonetics.md) |  | [optional] 
**Phrases** | Pointer to [**[]GetDictionaryLookup200ResponseEntryPhrasesInner**](GetDictionaryLookup200ResponseEntryPhrasesInner.md) |  | [optional] 
**Synonyms** | Pointer to [**[]GetDictionaryLookup200ResponseEntrySynonymsInner**](GetDictionaryLookup200ResponseEntrySynonymsInner.md) |  | [optional] 
**Word** | Pointer to **string** |  | [optional] 
**WordForms** | Pointer to [**[]GetDictionaryLookup200ResponseEntryWordFormsInner**](GetDictionaryLookup200ResponseEntryWordFormsInner.md) |  | [optional] 

## Methods

### NewGetDictionaryLookup200ResponseEntry

`func NewGetDictionaryLookup200ResponseEntry() *GetDictionaryLookup200ResponseEntry`

NewGetDictionaryLookup200ResponseEntry instantiates a new GetDictionaryLookup200ResponseEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDictionaryLookup200ResponseEntryWithDefaults

`func NewGetDictionaryLookup200ResponseEntryWithDefaults() *GetDictionaryLookup200ResponseEntry`

NewGetDictionaryLookup200ResponseEntryWithDefaults instantiates a new GetDictionaryLookup200ResponseEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefinitions

`func (o *GetDictionaryLookup200ResponseEntry) GetDefinitions() []GetDictionaryLookup200ResponseEntryDefinitionsInner`

GetDefinitions returns the Definitions field if non-nil, zero value otherwise.

### GetDefinitionsOk

`func (o *GetDictionaryLookup200ResponseEntry) GetDefinitionsOk() (*[]GetDictionaryLookup200ResponseEntryDefinitionsInner, bool)`

GetDefinitionsOk returns a tuple with the Definitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitions

`func (o *GetDictionaryLookup200ResponseEntry) SetDefinitions(v []GetDictionaryLookup200ResponseEntryDefinitionsInner)`

SetDefinitions sets Definitions field to given value.

### HasDefinitions

`func (o *GetDictionaryLookup200ResponseEntry) HasDefinitions() bool`

HasDefinitions returns a boolean if a field has been set.

### GetEnglishDefinitions

`func (o *GetDictionaryLookup200ResponseEntry) GetEnglishDefinitions() []GetDictionaryLookup200ResponseEntryEnglishDefinitionsInner`

GetEnglishDefinitions returns the EnglishDefinitions field if non-nil, zero value otherwise.

### GetEnglishDefinitionsOk

`func (o *GetDictionaryLookup200ResponseEntry) GetEnglishDefinitionsOk() (*[]GetDictionaryLookup200ResponseEntryEnglishDefinitionsInner, bool)`

GetEnglishDefinitionsOk returns a tuple with the EnglishDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnglishDefinitions

`func (o *GetDictionaryLookup200ResponseEntry) SetEnglishDefinitions(v []GetDictionaryLookup200ResponseEntryEnglishDefinitionsInner)`

SetEnglishDefinitions sets EnglishDefinitions field to given value.

### HasEnglishDefinitions

`func (o *GetDictionaryLookup200ResponseEntry) HasEnglishDefinitions() bool`

HasEnglishDefinitions returns a boolean if a field has been set.

### GetExamTags

`func (o *GetDictionaryLookup200ResponseEntry) GetExamTags() []string`

GetExamTags returns the ExamTags field if non-nil, zero value otherwise.

### GetExamTagsOk

`func (o *GetDictionaryLookup200ResponseEntry) GetExamTagsOk() (*[]string, bool)`

GetExamTagsOk returns a tuple with the ExamTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExamTags

`func (o *GetDictionaryLookup200ResponseEntry) SetExamTags(v []string)`

SetExamTags sets ExamTags field to given value.

### HasExamTags

`func (o *GetDictionaryLookup200ResponseEntry) HasExamTags() bool`

HasExamTags returns a boolean if a field has been set.

### GetExamples

`func (o *GetDictionaryLookup200ResponseEntry) GetExamples() []GetDictionaryLookup200ResponseEntryExamplesInner`

GetExamples returns the Examples field if non-nil, zero value otherwise.

### GetExamplesOk

`func (o *GetDictionaryLookup200ResponseEntry) GetExamplesOk() (*[]GetDictionaryLookup200ResponseEntryExamplesInner, bool)`

GetExamplesOk returns a tuple with the Examples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExamples

`func (o *GetDictionaryLookup200ResponseEntry) SetExamples(v []GetDictionaryLookup200ResponseEntryExamplesInner)`

SetExamples sets Examples field to given value.

### HasExamples

`func (o *GetDictionaryLookup200ResponseEntry) HasExamples() bool`

HasExamples returns a boolean if a field has been set.

### GetLanguage

`func (o *GetDictionaryLookup200ResponseEntry) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *GetDictionaryLookup200ResponseEntry) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *GetDictionaryLookup200ResponseEntry) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *GetDictionaryLookup200ResponseEntry) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetPhonetics

`func (o *GetDictionaryLookup200ResponseEntry) GetPhonetics() GetDictionaryLookup200ResponseEntryPhonetics`

GetPhonetics returns the Phonetics field if non-nil, zero value otherwise.

### GetPhoneticsOk

`func (o *GetDictionaryLookup200ResponseEntry) GetPhoneticsOk() (*GetDictionaryLookup200ResponseEntryPhonetics, bool)`

GetPhoneticsOk returns a tuple with the Phonetics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhonetics

`func (o *GetDictionaryLookup200ResponseEntry) SetPhonetics(v GetDictionaryLookup200ResponseEntryPhonetics)`

SetPhonetics sets Phonetics field to given value.

### HasPhonetics

`func (o *GetDictionaryLookup200ResponseEntry) HasPhonetics() bool`

HasPhonetics returns a boolean if a field has been set.

### GetPhrases

`func (o *GetDictionaryLookup200ResponseEntry) GetPhrases() []GetDictionaryLookup200ResponseEntryPhrasesInner`

GetPhrases returns the Phrases field if non-nil, zero value otherwise.

### GetPhrasesOk

`func (o *GetDictionaryLookup200ResponseEntry) GetPhrasesOk() (*[]GetDictionaryLookup200ResponseEntryPhrasesInner, bool)`

GetPhrasesOk returns a tuple with the Phrases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhrases

`func (o *GetDictionaryLookup200ResponseEntry) SetPhrases(v []GetDictionaryLookup200ResponseEntryPhrasesInner)`

SetPhrases sets Phrases field to given value.

### HasPhrases

`func (o *GetDictionaryLookup200ResponseEntry) HasPhrases() bool`

HasPhrases returns a boolean if a field has been set.

### GetSynonyms

`func (o *GetDictionaryLookup200ResponseEntry) GetSynonyms() []GetDictionaryLookup200ResponseEntrySynonymsInner`

GetSynonyms returns the Synonyms field if non-nil, zero value otherwise.

### GetSynonymsOk

`func (o *GetDictionaryLookup200ResponseEntry) GetSynonymsOk() (*[]GetDictionaryLookup200ResponseEntrySynonymsInner, bool)`

GetSynonymsOk returns a tuple with the Synonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynonyms

`func (o *GetDictionaryLookup200ResponseEntry) SetSynonyms(v []GetDictionaryLookup200ResponseEntrySynonymsInner)`

SetSynonyms sets Synonyms field to given value.

### HasSynonyms

`func (o *GetDictionaryLookup200ResponseEntry) HasSynonyms() bool`

HasSynonyms returns a boolean if a field has been set.

### GetWord

`func (o *GetDictionaryLookup200ResponseEntry) GetWord() string`

GetWord returns the Word field if non-nil, zero value otherwise.

### GetWordOk

`func (o *GetDictionaryLookup200ResponseEntry) GetWordOk() (*string, bool)`

GetWordOk returns a tuple with the Word field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWord

`func (o *GetDictionaryLookup200ResponseEntry) SetWord(v string)`

SetWord sets Word field to given value.

### HasWord

`func (o *GetDictionaryLookup200ResponseEntry) HasWord() bool`

HasWord returns a boolean if a field has been set.

### GetWordForms

`func (o *GetDictionaryLookup200ResponseEntry) GetWordForms() []GetDictionaryLookup200ResponseEntryWordFormsInner`

GetWordForms returns the WordForms field if non-nil, zero value otherwise.

### GetWordFormsOk

`func (o *GetDictionaryLookup200ResponseEntry) GetWordFormsOk() (*[]GetDictionaryLookup200ResponseEntryWordFormsInner, bool)`

GetWordFormsOk returns a tuple with the WordForms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWordForms

`func (o *GetDictionaryLookup200ResponseEntry) SetWordForms(v []GetDictionaryLookup200ResponseEntryWordFormsInner)`

SetWordForms sets WordForms field to given value.

### HasWordForms

`func (o *GetDictionaryLookup200ResponseEntry) HasWordForms() bool`

HasWordForms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


