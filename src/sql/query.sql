SELECT
  jets.name AS jetName,
  jets.age AS jetAge,
  jets.color AS jetColor,
  pilots.name AS pilotName,
  languages.language
FROM jets
JOIN pilots ON pilots.id = jets.pilot_id
LEFT JOIN pilot_languages ON pilot_languages.pilot_id = jets.pilot_id
LEFT JOIN languages ON languages.id = pilot_languages.language_id
WHERE TRUE
  {{ if ne .Age 0 -}}
  AND jets.age = :age
  {{- end}}
  {{if ne .PilotName "" -}}
  AND pilots.name LIKE :pilot_name
  {{- end}}
  {{if ne .JetName "" -}}
  AND jets.name LIKE :jet_name
  {{- end}}
  {{if ne .Language "" -}}
  AND languages.language = :language
  {{- end}}
ORDER BY jets.age, jets.id
;
