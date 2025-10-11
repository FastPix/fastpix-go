# BadRequest

Bad Request – Stream is either already enabled or cannot be enabled on trial plan.


## Supported Types

### TrialPlanRestrictionError

```go
badRequest := apierrors.CreateBadRequestTrialPlanRestrictionError(apierrors.TrialPlanRestrictionError{/* values here */})
```

### StreamAlreadyEnabledError

```go
badRequest := apierrors.CreateBadRequestStreamAlreadyEnabledError(apierrors.StreamAlreadyEnabledError{/* values here */})
```

