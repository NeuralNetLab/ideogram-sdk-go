# Describe

Response Types:

- <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#DescribeNewResponse">DescribeNewResponse</a>

Methods:

- <code title="post /describe">client.Describe.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#DescribeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#DescribeNewParams">DescribeNewParams</a>) (<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#DescribeNewResponse">DescribeNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Edit

Params Types:

- <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#MagicPromptOption">MagicPromptOption</a>
- <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ModelEnum">ModelEnum</a>
- <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#StyleTypeV2AndAbove">StyleTypeV2AndAbove</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#GenerateImage">GenerateImage</a>
- <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#StyleTypeV2AndAbove">StyleTypeV2AndAbove</a>

Methods:

- <code title="post /edit">client.Edit.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#EditService.Apply">Apply</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#EditApplyParams">EditApplyParams</a>) (<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#GenerateImage">GenerateImage</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Generate

Params Types:

- <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ImageRequestParam">ImageRequestParam</a>
- <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#MagicPromptVersionEnum">MagicPromptVersionEnum</a>

Methods:

- <code title="post /generate">client.Generate.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#GenerateService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#GenerateNewParams">GenerateNewParams</a>) (<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#GenerateImage">GenerateImage</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# InternalTesting

Params Types:

- <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#NestedObjectParam">NestedObjectParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#InternalTestingNewResponse">InternalTestingNewResponse</a>

Methods:

- <code title="post /internal-testing">client.InternalTesting.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#InternalTestingService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#InternalTestingNewParams">InternalTestingNewParams</a>) (<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#InternalTestingNewResponse">InternalTestingNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Manage

## API

Params Types:

- <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#PriceParam">PriceParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#Price">Price</a>
- <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#RechargeSettingsResponse">RechargeSettingsResponse</a>
- <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ManageAPIAddCreditsResponse">ManageAPIAddCreditsResponse</a>
- <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ManageAPIReactivateResponse">ManageAPIReactivateResponse</a>
- <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ManageAPIGetStripeSubscriptionResponse">ManageAPIGetStripeSubscriptionResponse</a>

Methods:

- <code title="post /manage/api/add_credits">client.Manage.API.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ManageAPIService.AddCredits">AddCredits</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ManageAPIAddCreditsParams">ManageAPIAddCreditsParams</a>) (<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ManageAPIAddCreditsResponse">ManageAPIAddCreditsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /manage/api/reactivate">client.Manage.API.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ManageAPIService.Reactivate">Reactivate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ManageAPIReactivateResponse">ManageAPIReactivateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /manage/api/stripe_subscription">client.Manage.API.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ManageAPIService.GetStripeSubscription">GetStripeSubscription</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ManageAPIGetStripeSubscriptionParams">ManageAPIGetStripeSubscriptionParams</a>) (<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ManageAPIGetStripeSubscriptionResponse">ManageAPIGetStripeSubscriptionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### APIKeys

Response Types:

- <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ManageApiapiKeyNewResponse">ManageApiapiKeyNewResponse</a>
- <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ManageApiapiKeyGetResponse">ManageApiapiKeyGetResponse</a>

Methods:

- <code title="post /manage/api/api_keys">client.Manage.API.APIKeys.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ManageAPIAPIKeyService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ManageApiapiKeyNewResponse">ManageApiapiKeyNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /manage/api/api_keys">client.Manage.API.APIKeys.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ManageAPIAPIKeyService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ManageApiapiKeyGetResponse">ManageApiapiKeyGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /manage/api/api_keys/{api_key_id}">client.Manage.API.APIKeys.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ManageAPIAPIKeyService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, apiKeyID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

### Subscription

Params Types:

- <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#RechargeSettingsSubscriptionParam">RechargeSettingsSubscriptionParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#RechargeSettingsSubscription">RechargeSettingsSubscription</a>
- <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ManageAPISubscriptionGetResponse">ManageAPISubscriptionGetResponse</a>
- <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ManageAPISubscriptionUpdateResponse">ManageAPISubscriptionUpdateResponse</a>

Methods:

- <code title="get /manage/api/subscription">client.Manage.API.Subscription.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ManageAPISubscriptionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ManageAPISubscriptionGetResponse">ManageAPISubscriptionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /manage/api/subscription">client.Manage.API.Subscription.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ManageAPISubscriptionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ManageAPISubscriptionUpdateParams">ManageAPISubscriptionUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ManageAPISubscriptionUpdateResponse">ManageAPISubscriptionUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Terms

Response Types:

- <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ManageAPITermGetResponse">ManageAPITermGetResponse</a>

Methods:

- <code title="get /manage/api/terms">client.Manage.API.Terms.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ManageAPITermService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ManageAPITermGetResponse">ManageAPITermGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /manage/api/terms">client.Manage.API.Terms.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ManageAPITermService.Accept">Accept</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ManageAPITermAcceptParams">ManageAPITermAcceptParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Reframe

Params Types:

- <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ResolutionImageGeneration">ResolutionImageGeneration</a>

Methods:

- <code title="post /reframe">client.Reframe.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ReframeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ReframeNewParams">ReframeNewParams</a>) (<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#GenerateImage">GenerateImage</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Remix

Methods:

- <code title="post /remix">client.Remix.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#RemixService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#RemixNewParams">RemixNewParams</a>) (<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#GenerateImage">GenerateImage</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Upscale

Methods:

- <code title="post /upscale">client.Upscale.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#UpscaleService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#UpscaleNewParams">UpscaleNewParams</a>) (<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#GenerateImage">GenerateImage</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# IdeogramV3

Params Types:

- <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#AspectRatio">AspectRatio</a>
- <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ColorPaletteUnionParam">ColorPaletteUnionParam</a>
- <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#RenderingSpeed">RenderingSpeed</a>
- <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ResolutionIdeogram">ResolutionIdeogram</a>
- <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#StyleTypeV3">StyleTypeV3</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ImageGenerationResponse">ImageGenerationResponse</a>
- <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ResolutionIdeogram">ResolutionIdeogram</a>
- <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#StyleTypeV3">StyleTypeV3</a>

Methods:

- <code title="post /v1/ideogram-v3/edit">client.IdeogramV3.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#IdeogramV3Service.Edit">Edit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#IdeogramV3EditParams">IdeogramV3EditParams</a>) (<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ImageGenerationResponse">ImageGenerationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/ideogram-v3/generate">client.IdeogramV3.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#IdeogramV3Service.Generate">Generate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#IdeogramV3GenerateParams">IdeogramV3GenerateParams</a>) (<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ImageGenerationResponse">ImageGenerationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/ideogram-v3/reframe">client.IdeogramV3.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#IdeogramV3Service.Reframe">Reframe</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#IdeogramV3ReframeParams">IdeogramV3ReframeParams</a>) (<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ImageGenerationResponse">ImageGenerationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/ideogram-v3/remix">client.IdeogramV3.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#IdeogramV3Service.Remix">Remix</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#IdeogramV3RemixParams">IdeogramV3RemixParams</a>) (<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ImageGenerationResponse">ImageGenerationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/ideogram-v3/replace-background">client.IdeogramV3.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#IdeogramV3Service.ReplaceBackground">ReplaceBackground</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#IdeogramV3ReplaceBackgroundParams">IdeogramV3ReplaceBackgroundParams</a>) (<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go">ideogramsdk</a>.<a href="https://pkg.go.dev/github.com/NeuralNetLab/ideogram-sdk-go#ImageGenerationResponse">ImageGenerationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
