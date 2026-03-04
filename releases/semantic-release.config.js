/**
 * Semantic Release Configuration
 * Automated versioning and changelog generation
 * Mode: Full AI Autonomy (auto-approve all releases)
 */

module.exports = {
  branches: [
    'main',
    { name: 'develop', prerelease: 'alpha' },
  ],
  plugins: [
    // Analyze commits using conventional-changelog
    ['@semantic-release/commit-analyzer', {
      preset: 'conventionalcommits',
      releaseRules: [
        { breaking: true, release: 'major' },
        { type: 'feat', release: 'minor' },
        { type: 'fix', release: 'patch' },
        { type: 'perf', release: 'patch' },
        { type: 'refactor', release: false },
        { type: 'docs', release: false },
        { type: 'test', release: false },
        { type: 'chore', release: false },
      ],
    }],

    // Generate release notes
    ['@semantic-release/release-notes-generator', {
      preset: 'conventionalcommits',
      presetConfig: {
        types: [
          { type: 'feat', section: 'Features' },
          { type: 'fix', section: 'Bug Fixes' },
          { type: 'perf', section: 'Performance Improvements' },
          { type: 'revert', section: 'Reverts' },
          { type: 'refactor', section: 'Code Refactoring', hidden: true },
          { type: 'docs', section: 'Documentation', hidden: true },
          { type: 'test', section: 'Tests', hidden: true },
          { type: 'chore', section: 'Chores', hidden: true },
        ],
      },
    }],

    // Update changelog
    ['@semantic-release/changelog', {
      changelogFile: 'CHANGELOG.md',
    }],

    // Update version file (Go doesn't use package.json)
    ['@semantic-release/exec', {
      prepareCmd: 'echo "${nextRelease.version}" > VERSION',
    }],

    // Create GitHub release
    ['@semantic-release/github', {
      assets: [
        { path: 'CHANGELOG.md', label: 'Changelog' },
      ],
    }],

    // Commit release artifacts
    ['@semantic-release/git', {
      assets: ['CHANGELOG.md', 'VERSION'],
      message: 'chore(release): ${nextRelease.version} [skip ci]\n\n${nextRelease.notes}',
    }],
  ],
};
