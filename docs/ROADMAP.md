```mermaid
---
displayMode: compact
---
gantt
  dateFormat YYYY-MM-DD
  title Citadel Magic GANTT diagram
  excludes weekends

  Initial milestone: milestone, m_initial, 2024-09-06, 0d
  V1.0.0-alpha.1: milestone, m1, after validation_cmd, 0d
  V1.0.0-alpha.2: milestone, m2, after generate_cmd dry_mode, 0d
  V1.0.0: milestone, m3, after diff_mode cycle_detection, 0d

  Section Comands
  Validation Command: validation_cmd, after manifest_validation, 1d
  Generate Command: generate_cmd, after manifest_validation Magic_apply, 1d

  Section Modes
  Dry Mode: dry_mode, after Magic_apply, 1d
  Magic Apply: Magic_apply, after Magic_render, 2d
  Diff Mode: diff_mode, after overwrite_mode, 3d
  Overwrite Mode: overwrite_mode, after clean_mode, 1d
  Clean Mode: clean_mode, after Magic_apply, 1d

  Section Rendering
  Magic Render: Magic_render, after Cast_render, 1d
  Cast Render: Cast_render, after template_source, 1d
  Conditional Render: conditional_render, after m1, 2d
  Collection Render: collection_render, after m1, 3d
  Conditional Collection Render: conditional_collection_render, after collection_render, 2d
  
  Section Validation
  Manifest process: manifest_process, after m_initial, 1d
  Manifest Schema Validation: manifest_validation, after manifest_process, 1d
  Cycle Detection: cycle_detection, after Magic_source, 2d

  Section Sources
  Template Source: template_source, after m_initial, 1d
  File Source: file_source, after template_source, 1d
  Document Source: document_source, after file_source, 2d
  Magic Source: Magic_source, after Magic_render, 5d
  Cast Selection: Cast_selection, after Magic_source, 2d
```