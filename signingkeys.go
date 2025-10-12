package fastpixgo

import (
	"bytes"
	"context"
	"fmt"
	"github.com/FastPix/fastpix-go/internal/config"
	"github.com/FastPix/fastpix-go/internal/hooks"
	"github.com/FastPix/fastpix-go/internal/utils"
	"github.com/FastPix/fastpix-go/models/apierrors"
	"github.com/FastPix/fastpix-go/models/components"
	"github.com/FastPix/fastpix-go/models/operations"
	"github.com/FastPix/fastpix-go/retry"
	"net/http"
	"net/url"
)

type SigningKeys struct {
	rootSDK          *Fastpixgo
	sdkConfiguration config.SDKConfiguration
	hooks            *hooks.Hooks
}

func newSigningKeys(rootSDK *Fastpixgo, sdkConfig config.SDKConfiguration, hooks *hooks.Hooks) *SigningKeys {
	return &SigningKeys{
		rootSDK:          rootSDK,
		sdkConfiguration: sdkConfig,
		hooks:            hooks,
	}
}

// CreateSigningKey - Create a signing key
// This endpoint allows you to create a new signing key pair for FastPix. When you call this endpoint, the API generates a 2048-bit RSA key pair. The privateKey will be returned in the response, encoded in Base64 format, and you will receive a unique key id to reference the key in future operations. FastPix will securely store the public key to validate signed tokens.
//
// <h4>Instructions</h4>
//
// **Private key handling:** The privateKey you receive is encoded in Base64. To use it, you'll need to decode it using Base64 decoding. Make sure to store this private key securely, as it is required for signing tokens.
//
// **Key-ID:** The id will be used to reference this specific key pair in future API requests or configurations.
//
// Once the key pair is generated, the private key must be securely stored by the developer, as FastPix will not save it. The public key will be used by FastPix to verify any signed tokens, ensuring that the client interacting with the system is legitimate.
//
// <h4>Use case scenario</h4>
//
// **Use case:** A developer building a video subscription service wants to ensure that only authorized users can access premium content. By generating a signing key, the developer can issue signed JSON Web Tokens (JWTs) to authenticate and authorize users. These tokens can be validated by FastPix using the stored public key.
//
// **Detailed example:**  Imagine a scenario where you're building a video-on-demand platform that restricts access based on user subscriptions. To ensure only subscribed users can stream content, you generate a signing key using this API. Each time a user logs in, you create a JWT signed with the private key. When the user attempts to play a video, FastPix uses the public key to verify the token and confirms that the user is authorized.
func (s *SigningKeys) CreateSigningKey(ctx context.Context, opts ...operations.Option) (*operations.CreateSigningKeyResponse, error) {
	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionRetries,
		operations.SupportedOptionTimeout,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}

	var baseURL string
	if o.ServerURL == nil {
		baseURL = utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	} else {
		baseURL = *o.ServerURL
	}
	opURL, err := url.JoinPath(baseURL, "/iam/signing-keys")
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	hookCtx := hooks.HookContext{
		SDK:              s.rootSDK,
		SDKConfiguration: s.sdkConfiguration,
		BaseURL:          baseURL,
		Context:          ctx,
		OperationID:      "create_signing_key",
		SecuritySource:   s.sdkConfiguration.Security,
	}

	timeout := o.Timeout
	if timeout == nil {
		timeout = s.sdkConfiguration.Timeout
	}

	if timeout != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *timeout)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, "POST", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	for k, v := range o.SetHeaders {
		req.Header.Set(k, v)
	}

	globalRetryConfig := s.sdkConfiguration.RetryConfig
	retryConfig := o.Retries
	if retryConfig == nil {
		if globalRetryConfig != nil {
			retryConfig = globalRetryConfig
		} else {
			retryConfig = &retry.Config{
				Strategy: "backoff", Backoff: &retry.BackoffStrategy{
					InitialInterval: 1000,
					MaxInterval:     10000,
					Exponent:        1.5,
					MaxElapsedTime:  3600000,
				},
				RetryConnectionErrors: true,
			}
		}
	}

	var httpRes *http.Response
	if retryConfig != nil {
		httpRes, err = utils.Retry(ctx, utils.Retries{
			Config: retryConfig,
			StatusCodes: []string{
				"408",
				"429",
				"500",
				"502",
				"503",
				"504",
			},
		}, func() (*http.Response, error) {
			if req.Body != nil && req.Body != http.NoBody && req.GetBody != nil {
				copyBody, err := req.GetBody()

				if err != nil {
					return nil, err
				}

				req.Body = copyBody
			}

			req, err = s.hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
			if err != nil {
				if retry.IsPermanentError(err) || retry.IsTemporaryError(err) {
					return nil, err
				}

				return nil, retry.Permanent(err)
			}

			httpRes, err := s.sdkConfiguration.Client.Do(req)
			if err != nil || httpRes == nil {
				if err != nil {
					err = fmt.Errorf("error sending request: %w", err)
				} else {
					err = fmt.Errorf("error sending request: no response")
				}

				_, err = s.hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			}
			return httpRes, err
		})

		if err != nil {
			return nil, err
		} else {
			httpRes, err = s.hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	} else {
		req, err = s.hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
		if err != nil {
			return nil, err
		}

		httpRes, err = s.sdkConfiguration.Client.Do(req)
		if err != nil || httpRes == nil {
			if err != nil {
				err = fmt.Errorf("error sending request: %w", err)
			} else {
				err = fmt.Errorf("error sending request: no response")
			}

			_, err = s.hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			return nil, err
		} else if utils.MatchStatusCodes([]string{"401", "403", "4XX", "5XX"}, httpRes.StatusCode) {
			_httpRes, err := s.hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
			if err != nil {
				return nil, err
			} else if _httpRes != nil {
				httpRes = _httpRes
			}
		} else {
			httpRes, err = s.hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	}

	res := &operations.CreateSigningKeyResponse{
		HTTPMeta: components.HTTPMetadata{
			Request:  req,
			Response: httpRes,
		},
	}

	switch {
	case httpRes.StatusCode == 201:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out components.CreateResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.CreateResponse = &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.UnAuthorizedResponseError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			out.HTTPMeta = components.HTTPMetadata{
				Request:  req,
				Response: httpRes,
			}
			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 403:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.ForbiddenResponseError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			out.HTTPMeta = components.HTTPMetadata{
				Request:  req,
				Response: httpRes,
			}
			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil

}

// ListSigningKeys - Get list of signing key
// This endpoint returns a list of all the signing keys associated with an organization in FastPix. Each key entry in the response includes metadata such as the key id, creation date, and workspace details. This helps you manage multiple keys, track their usage, and identify which keys are valid for signing API requests.
//
// <h4>How it works</h4>
//
// The API returns the list in a paginated format, allowing you to audit and track all keys used for your application. Regularly reviewing this list is essential for ensuring that old or compromised keys are promptly revoked and that new keys are properly integrated into workflows.
//
// <h4>Use case scenario</h4>
//
// **Use case:** A security-conscious development team wants to ensure they follow a key rotation policy, rotating signing keys every few months. By retrieving the list of signing keys, they can identify which keys are still in use and which ones need to be rotated.
//
// **Detailed example:**  You're managing a multi-region video platform where different teams in different regions have created their own signing keys. To ensure compliance with your organization's security policies, you regularly review the list of signing keys to verify which ones are still active. You find a few keys that haven’t been used in months, and based on the creation date, you decide to rotate them.
func (s *SigningKeys) ListSigningKeys(ctx context.Context, limit *float64, offset *float64, opts ...operations.Option) (*operations.ListSigningKeysResponse, error) {
	request := operations.ListSigningKeysRequest{
		Limit:  limit,
		Offset: offset,
	}

	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionRetries,
		operations.SupportedOptionTimeout,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}

	var baseURL string
	if o.ServerURL == nil {
		baseURL = utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	} else {
		baseURL = *o.ServerURL
	}
	opURL, err := url.JoinPath(baseURL, "/iam/signing-keys")
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	hookCtx := hooks.HookContext{
		SDK:              s.rootSDK,
		SDKConfiguration: s.sdkConfiguration,
		BaseURL:          baseURL,
		Context:          ctx,
		OperationID:      "list_signing_keys",
		SecuritySource:   s.sdkConfiguration.Security,
	}

	timeout := o.Timeout
	if timeout == nil {
		timeout = s.sdkConfiguration.Timeout
	}

	if timeout != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *timeout)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, "GET", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	for k, v := range o.SetHeaders {
		req.Header.Set(k, v)
	}

	globalRetryConfig := s.sdkConfiguration.RetryConfig
	retryConfig := o.Retries
	if retryConfig == nil {
		if globalRetryConfig != nil {
			retryConfig = globalRetryConfig
		} else {
			retryConfig = &retry.Config{
				Strategy: "backoff", Backoff: &retry.BackoffStrategy{
					InitialInterval: 1000,
					MaxInterval:     10000,
					Exponent:        1.5,
					MaxElapsedTime:  3600000,
				},
				RetryConnectionErrors: true,
			}
		}
	}

	var httpRes *http.Response
	if retryConfig != nil {
		httpRes, err = utils.Retry(ctx, utils.Retries{
			Config: retryConfig,
			StatusCodes: []string{
				"408",
				"429",
				"500",
				"502",
				"503",
				"504",
			},
		}, func() (*http.Response, error) {
			if req.Body != nil && req.Body != http.NoBody && req.GetBody != nil {
				copyBody, err := req.GetBody()

				if err != nil {
					return nil, err
				}

				req.Body = copyBody
			}

			req, err = s.hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
			if err != nil {
				if retry.IsPermanentError(err) || retry.IsTemporaryError(err) {
					return nil, err
				}

				return nil, retry.Permanent(err)
			}

			httpRes, err := s.sdkConfiguration.Client.Do(req)
			if err != nil || httpRes == nil {
				if err != nil {
					err = fmt.Errorf("error sending request: %w", err)
				} else {
					err = fmt.Errorf("error sending request: no response")
				}

				_, err = s.hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			}
			return httpRes, err
		})

		if err != nil {
			return nil, err
		} else {
			httpRes, err = s.hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	} else {
		req, err = s.hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
		if err != nil {
			return nil, err
		}

		httpRes, err = s.sdkConfiguration.Client.Do(req)
		if err != nil || httpRes == nil {
			if err != nil {
				err = fmt.Errorf("error sending request: %w", err)
			} else {
				err = fmt.Errorf("error sending request: no response")
			}

			_, err = s.hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			return nil, err
		} else if utils.MatchStatusCodes([]string{"401", "403", "422", "4XX", "5XX"}, httpRes.StatusCode) {
			_httpRes, err := s.hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
			if err != nil {
				return nil, err
			} else if _httpRes != nil {
				httpRes = _httpRes
			}
		} else {
			httpRes, err = s.hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	}

	res := &operations.ListSigningKeysResponse{
		HTTPMeta: components.HTTPMetadata{
			Request:  req,
			Response: httpRes,
		},
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out components.GetAllSigningKeyResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.GetAllSigningKeyResponse = &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.UnAuthorizedResponseError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			out.HTTPMeta = components.HTTPMetadata{
				Request:  req,
				Response: httpRes,
			}
			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 403:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.ForbiddenResponseError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			out.HTTPMeta = components.HTTPMetadata{
				Request:  req,
				Response: httpRes,
			}
			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.ValidationErrorResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			out.HTTPMeta = components.HTTPMetadata{
				Request:  req,
				Response: httpRes,
			}
			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil

}

// DeleteSigningKey - Delete a signing key
// This endpoint allows you to delete an existing signing key, and the action is permanent. Once a key is deleted, any signatures or tokens generated using that key will immediately become invalid. This means you can no longer use the key to sign JSON Web Tokens (JWTs) or authenticate API requests.
// <h4>Usage</h4>
// To delete a signing key, you will need to provide the unique key id that was obtained when creating the signing key. This key id serves as the identifier for the specific signing key you want to remove from your account.
//
// <h4>How it works</h4>
//
// By specifying the key id, the API removes the signing key from the system. After the key is deleted, any API requests or tokens that rely on it will fail. This action is useful when a key is compromised or when rotating keys as part of security policies.
//
// <h4>Use case scenario</h4>
//
// **Use case:** A key used by an outdated application version has been compromised, or a developer accidentally leaked it. To prevent unauthorized access, the developer deletes the signing key, revoking its ability to sign requests immediately.
//
// **Detailed example:**  Let’s say you have a signing key used for a specific version of your mobile app, and you discover that this key has been compromised due to a security breach. To mitigate the issue, you delete the key to invalidate any tokens generated using it. As soon as the key is deleted, users on the compromised version of the app can no longer make valid requests, thus preventing further exploitation.
func (s *SigningKeys) DeleteSigningKey(ctx context.Context, signingKeyID string, opts ...operations.Option) (*operations.DeleteSigningKeyResponse, error) {
	request := operations.DeleteSigningKeyRequest{
		SigningKeyID: signingKeyID,
	}

	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionRetries,
		operations.SupportedOptionTimeout,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}

	var baseURL string
	if o.ServerURL == nil {
		baseURL = utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	} else {
		baseURL = *o.ServerURL
	}
	opURL, err := utils.GenerateURL(ctx, baseURL, "/iam/signing-keys/{signingKeyId}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	hookCtx := hooks.HookContext{
		SDK:              s.rootSDK,
		SDKConfiguration: s.sdkConfiguration,
		BaseURL:          baseURL,
		Context:          ctx,
		OperationID:      "delete_signing_key",
		SecuritySource:   s.sdkConfiguration.Security,
	}

	timeout := o.Timeout
	if timeout == nil {
		timeout = s.sdkConfiguration.Timeout
	}

	if timeout != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *timeout)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, "DELETE", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	for k, v := range o.SetHeaders {
		req.Header.Set(k, v)
	}

	globalRetryConfig := s.sdkConfiguration.RetryConfig
	retryConfig := o.Retries
	if retryConfig == nil {
		if globalRetryConfig != nil {
			retryConfig = globalRetryConfig
		} else {
			retryConfig = &retry.Config{
				Strategy: "backoff", Backoff: &retry.BackoffStrategy{
					InitialInterval: 1000,
					MaxInterval:     10000,
					Exponent:        1.5,
					MaxElapsedTime:  3600000,
				},
				RetryConnectionErrors: true,
			}
		}
	}

	var httpRes *http.Response
	if retryConfig != nil {
		httpRes, err = utils.Retry(ctx, utils.Retries{
			Config: retryConfig,
			StatusCodes: []string{
				"408",
				"429",
				"500",
				"502",
				"503",
				"504",
			},
		}, func() (*http.Response, error) {
			if req.Body != nil && req.Body != http.NoBody && req.GetBody != nil {
				copyBody, err := req.GetBody()

				if err != nil {
					return nil, err
				}

				req.Body = copyBody
			}

			req, err = s.hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
			if err != nil {
				if retry.IsPermanentError(err) || retry.IsTemporaryError(err) {
					return nil, err
				}

				return nil, retry.Permanent(err)
			}

			httpRes, err := s.sdkConfiguration.Client.Do(req)
			if err != nil || httpRes == nil {
				if err != nil {
					err = fmt.Errorf("error sending request: %w", err)
				} else {
					err = fmt.Errorf("error sending request: no response")
				}

				_, err = s.hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			}
			return httpRes, err
		})

		if err != nil {
			return nil, err
		} else {
			httpRes, err = s.hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	} else {
		req, err = s.hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
		if err != nil {
			return nil, err
		}

		httpRes, err = s.sdkConfiguration.Client.Do(req)
		if err != nil || httpRes == nil {
			if err != nil {
				err = fmt.Errorf("error sending request: %w", err)
			} else {
				err = fmt.Errorf("error sending request: no response")
			}

			_, err = s.hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			return nil, err
		} else if utils.MatchStatusCodes([]string{"401", "403", "404", "422", "4XX", "5XX"}, httpRes.StatusCode) {
			_httpRes, err := s.hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
			if err != nil {
				return nil, err
			} else if _httpRes != nil {
				httpRes = _httpRes
			}
		} else {
			httpRes, err = s.hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	}

	res := &operations.DeleteSigningKeyResponse{
		HTTPMeta: components.HTTPMetadata{
			Request:  req,
			Response: httpRes,
		},
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out components.DeleteSigningKeyResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.DeleteSigningKeyResponse = &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.UnAuthorizedResponseError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			out.HTTPMeta = components.HTTPMetadata{
				Request:  req,
				Response: httpRes,
			}
			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 403:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.ForbiddenResponseError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			out.HTTPMeta = components.HTTPMetadata{
				Request:  req,
				Response: httpRes,
			}
			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.SigningKeyNotFoundError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			out.HTTPMeta = components.HTTPMetadata{
				Request:  req,
				Response: httpRes,
			}
			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.ValidationErrorResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			out.HTTPMeta = components.HTTPMetadata{
				Request:  req,
				Response: httpRes,
			}
			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil

}

// GetSigningKeyByID - Get signing key by ID
// This endpoint allows you to retrieve detailed information about a specific signing key using its unique key id. While the private key is not returned for security reasons, you'll be able to see the key's creation date, status, and other associated metadata. This endpoint also returns the workspaceId and publicKey in the response.
//
// <h4>Usage: Generating a JWT token</h4>
//
// In the response, you will receive the workspaceId and publicKey associated with the signing key. With the publicKey and the privateKey obtained from the "Create a Signing Key" endpoint, you can generate a JSON Web Token (JWT) using the RS256 algorithm. This token can be utilized for accessing private media assets, GIFs, thumbnails, and spritesheets.
//
// <h4>Payload:</h4>
//
// ```
// {
//
//	"kid": "359302ee-2446-4afe-9348-8b4656b9ddb1",
//	"aud": "media:6cee6f85-9334-4a51-9ce3-e0241d94ceef",
//	"iss": "fastpix.io",
//	"sub": "",
//	"iat": 1706703204,
//	"exp": 1735626783
//
// }
// ```
//
// * **kid:** The key ID of the signing key.
// * **aud:** The audience for which the token is intended.
// * **iss:** The issuer of the token (e.g., "fastpix.io").
// * **sub:** The subject of the token, typically representing the user or entity the token is issued for. In this case, use the workspaceId fetched from the "Get Signing Key by ID" endpoint.
// * **groups:** An array of groups the subject belongs to (e.g., ["user"]).
// * **iat:** The issued-at timestamp, indicating when the token was created.
// * **exp:** The expiration timestamp, indicating when the token will no longer be valid.
//
// <h4>Use case scenario</h4>
//
// **Use case:** A developer is unsure about the status of a signing key they created months ago and wants to verify whether it's still in use or has expired.
//
// **Detailed example:**  You’re working on a streaming platform and realize you haven’t checked the status of a signing key that was used for playback access several months ago. By fetching the key details using its ID, you can confirm whether it’s still active, when it was created, and if it’s nearing expiration. This allows you to plan a rotation or deactivation if needed.
func (s *SigningKeys) GetSigningKeyByID(ctx context.Context, signingKeyID string, opts ...operations.Option) (*operations.GetSigningKeyByIDResponse, error) {
	request := operations.GetSigningKeyByIDRequest{
		SigningKeyID: signingKeyID,
	}

	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionRetries,
		operations.SupportedOptionTimeout,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}

	var baseURL string
	if o.ServerURL == nil {
		baseURL = utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	} else {
		baseURL = *o.ServerURL
	}
	opURL, err := utils.GenerateURL(ctx, baseURL, "/iam/signing-keys/{signingKeyId}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	hookCtx := hooks.HookContext{
		SDK:              s.rootSDK,
		SDKConfiguration: s.sdkConfiguration,
		BaseURL:          baseURL,
		Context:          ctx,
		OperationID:      "get-signing_key_by_id",
		SecuritySource:   s.sdkConfiguration.Security,
	}

	timeout := o.Timeout
	if timeout == nil {
		timeout = s.sdkConfiguration.Timeout
	}

	if timeout != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *timeout)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, "GET", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	for k, v := range o.SetHeaders {
		req.Header.Set(k, v)
	}

	globalRetryConfig := s.sdkConfiguration.RetryConfig
	retryConfig := o.Retries
	if retryConfig == nil {
		if globalRetryConfig != nil {
			retryConfig = globalRetryConfig
		} else {
			retryConfig = &retry.Config{
				Strategy: "backoff", Backoff: &retry.BackoffStrategy{
					InitialInterval: 1000,
					MaxInterval:     10000,
					Exponent:        1.5,
					MaxElapsedTime:  3600000,
				},
				RetryConnectionErrors: true,
			}
		}
	}

	var httpRes *http.Response
	if retryConfig != nil {
		httpRes, err = utils.Retry(ctx, utils.Retries{
			Config: retryConfig,
			StatusCodes: []string{
				"408",
				"429",
				"500",
				"502",
				"503",
				"504",
			},
		}, func() (*http.Response, error) {
			if req.Body != nil && req.Body != http.NoBody && req.GetBody != nil {
				copyBody, err := req.GetBody()

				if err != nil {
					return nil, err
				}

				req.Body = copyBody
			}

			req, err = s.hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
			if err != nil {
				if retry.IsPermanentError(err) || retry.IsTemporaryError(err) {
					return nil, err
				}

				return nil, retry.Permanent(err)
			}

			httpRes, err := s.sdkConfiguration.Client.Do(req)
			if err != nil || httpRes == nil {
				if err != nil {
					err = fmt.Errorf("error sending request: %w", err)
				} else {
					err = fmt.Errorf("error sending request: no response")
				}

				_, err = s.hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			}
			return httpRes, err
		})

		if err != nil {
			return nil, err
		} else {
			httpRes, err = s.hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	} else {
		req, err = s.hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
		if err != nil {
			return nil, err
		}

		httpRes, err = s.sdkConfiguration.Client.Do(req)
		if err != nil || httpRes == nil {
			if err != nil {
				err = fmt.Errorf("error sending request: %w", err)
			} else {
				err = fmt.Errorf("error sending request: no response")
			}

			_, err = s.hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
			return nil, err
		} else if utils.MatchStatusCodes([]string{"401", "403", "404", "422", "4XX", "5XX"}, httpRes.StatusCode) {
			_httpRes, err := s.hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
			if err != nil {
				return nil, err
			} else if _httpRes != nil {
				httpRes = _httpRes
			}
		} else {
			httpRes, err = s.hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
			if err != nil {
				return nil, err
			}
		}
	}

	res := &operations.GetSigningKeyByIDResponse{
		HTTPMeta: components.HTTPMetadata{
			Request:  req,
			Response: httpRes,
		},
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out components.GetPublicPemUsingSigningKeyIDResponseDTO
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.GetPublicPemUsingSigningKeyIDResponseDTO = &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 401:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.UnAuthorizedResponseError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			out.HTTPMeta = components.HTTPMetadata{
				Request:  req,
				Response: httpRes,
			}
			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 403:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.ForbiddenResponseError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			out.HTTPMeta = components.HTTPMetadata{
				Request:  req,
				Response: httpRes,
			}
			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.SigningKeyNotFoundError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			out.HTTPMeta = components.HTTPMetadata{
				Request:  req,
				Response: httpRes,
			}
			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out apierrors.ValidationErrorResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			out.HTTPMeta = components.HTTPMetadata{
				Request:  req,
				Response: httpRes,
			}
			return nil, &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, apierrors.NewAPIError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, apierrors.NewAPIError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil

}
