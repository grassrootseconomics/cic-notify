package template

import "bytes"

func (t *TxNotifyTemplates) Prepare(template TemplateType, templatePayload interface{}) (string, error) {
	var preparedTemplate bytes.Buffer

	switch template {
	case FailedTemeplate:
		if err := t.failedTmpl.Execute(&preparedTemplate, templatePayload); err != nil {
			return "", err
		}
	case SuccessReceivedTemplate:
		if err := t.successReceivedTmpl.Execute(&preparedTemplate, templatePayload); err != nil {
			return "", err
		}
	case SuccessSentTemplate:
		if err := t.successSentTmpl.Execute(&preparedTemplate, templatePayload); err != nil {
			return "", err
		}
	}

	return preparedTemplate.String(), nil
}
