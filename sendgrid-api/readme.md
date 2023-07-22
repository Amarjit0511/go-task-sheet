json response:

{
  "from_name": "Sender Name",
  "from_email": "amarjitkrxs@gmail.com",
  "to_name": "Recipient Name",
  "to_email": "amarjitkr0511@gmail.com",
  "subject": "Email Subject",
  "plain_text_content": "Plain text content of the email",
  "html_content": "<strong>HTML content of the email</strong>"
}


table:

CREATE TABLE sendgrid_response (
    id SERIAL PRIMARY KEY,
    status_code INT NOT NULL,
    body TEXT NOT NULL,
    headers JSONB NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

select * from sendgrid_response;

GRANT SELECT, INSERT, UPDATE, DELETE ON sendgrid_response TO amarjit;

GRANT USAGE ON SEQUENCE sendgrid_response_id_seq TO amarjit;
