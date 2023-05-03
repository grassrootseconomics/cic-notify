--name: create-at-receipt
-- Create a new AfricasTalking status receipt
-- $1: status_code
-- $2: message_id
INSERT INTO at_receipts(
    status_code,
    message_id,
) VALUES($1, $2) RETURNING id

--name: create-tg-receipt
-- Create a new Telegram status receipt
-- $2: message_id
INSERT INTO tg_receipts(
    message_id,
) VALUES($1) RETURNING id

--name: set-at-delivered
-- Mark an AfricasTalking SMS as successful
-- $1: message_id
UPDATE at_receipts SET delivered = true WHERE message_id=$1
