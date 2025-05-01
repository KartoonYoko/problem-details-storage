# problem-details-storage

### Entities
- backet - хранилище определённых ошибок, например, для определённого сервиса
  - id pk
  - name
  - description
- backet-problem-details
  - id pk
  - problem-details-id 
- problem-details
  - id pk
  - description string NOT NULL
  - type - string
  - title - string
  - status - int
  - detail - string
  - instance - string
- extension ?? как сделать extensions
  - key string
  - value jsonb
- label - метка для фильтрации ошибок
  - id pk
  - name string not null
- problem-details-label
  - id pk
  - problem-details-id
  - label-id