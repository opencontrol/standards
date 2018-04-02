# Standards

To import these data into a OpenControl project add the follow code to your opencontrol.yaml file.
```yaml
dependencies:
  standards:
    - url: https://github.com/opencontrol/standards
      revision: master
```

For more information on the opencontrol.yaml visit the [Compliance Masonry CLI](https://github.com/opencontrol/compliance-masonry#creating-an-opencontrol-project).


## Using the NIST 800-53 standards
On 1-APR-2018, the NIST 800-53 standard files were updated to allow OpenControl content authors to specify which revision of NIST 800-53 should be used. There are now three options:

- NIST-800-53 rev3: Legacy NIST 800-53 revision 3 content [[nist-800-53-rev3.yaml](https://github.com/opencontrol/standards/blob/master/nist-800-53-rev3.yaml)]
- NIST-800-53 rev4: NIST 800-53 revision 4 content [[nist-800-53-rev4.yaml](https://github.com/opencontrol/standards/blob/master/nist-800-53-rev3.yaml)]
- NIST-800-53: Special target that will always point to latest NIST revision upon ratification from NIST. Currently this is rev4, and will automatically  transition to rev5 when available. [[nist-800-53-latest.yaml](https://github.com/shawndwells/standards/blob/master/nist-800-53-latest.yaml)]

### Sample component.yaml
When creating OpenControl content for components, the ``standard_key`` should be updated to reflect which NIST 800-53 content should be used.

For example, for NIST 800-53 rev4:
`````
- control_key: AC-2 (2)
  standard_key: NIST-800-53 rev4
  covered_by: []
  implementation_status: Implemented
  narrative:
    - text: |
        'Your control response text'
`````

And for your content to automatically use the latest revision, use:

`````
- control_key: AC-2 (2)
  standard_key: NIST-800-53
  covered_by: []
  implementation_status: Not applicable
  narrative:
    - text: |
        'Your control response text'
`````

