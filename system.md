```mermaid
erDiagram
    HUMAN ||--|{ MOTIVATIONS: provides
    HUMAN ||--|| RESOURCES: provides
    DRIVER ||--|{ MOTIVATIONS : has
    DRIVER ||--|| RESOURCES : factors_in
    DRIVER ||--|| OPERATOR : gives_goal
    DRIVER ||--|| OVERALL_STATE: factors_in
    OPERATOR ||--|| RESOURCES : factors_in
    OPERATOR ||--|{ OPERATIONS : generates
    OPERATOR ||--|| OVERALL_STATE: factors_in
    ACTORS ||--|| RESOURCES : uses
    ACTORS }|--|| OPERATIONS : executes
    ACTORS ||--|{ OUTCOMES : creates
    INTERPRETERS ||--|{ OUTCOMES : interprets
    INTERPRETERS }|--|| OVERALL_STATE : updates

    HUMAN

    DRIVER {
        prompt what_would_you_do
    }

    OPERATOR {
        prompt what_would_you_do
        goal your_goal_is_to
    }

    ACTORS {
        prompt try_and_perform_these_operations
        operations order_of_operations
    }

    INTERPRETERS {

    }

    OUTCOMES {
        state_changes string

    }

    OPERATIONS {
        ordered_operations string_map
    }

    MOTIVATIONS {
        motivations string_map
    }

    RESOURCES {
        money number
        credentials map_string_map
    }

    OVERALL_STATE {
        current_state string
    }


