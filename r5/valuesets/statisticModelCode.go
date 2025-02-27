// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/statistic-model-code
*/type StatisticModelCode string

const (
	// Used for one-tailed test (1 threshold), no additional elements needed
	StatisticModelCodeOneTailedTest StatisticModelCode = "oneTailedTest"
	// Used for two-tailed test (2 threshold), no additional elements needed
	StatisticModelCodeTwoTailedTest StatisticModelCode = "twoTailedTest"
	// Used for z-test, no additional elements needed
	StatisticModelCodeZTest StatisticModelCode = "zTest"
	// Used for 1-sample t-test, may be paired with "value" to express degrees of freedom
	StatisticModelCodeOneSampleTTest StatisticModelCode = "oneSampleTTest"
	// Used for 2-sample t-test, may be paired with "value" to express degrees of freedom
	StatisticModelCodeTwoSampleTTest StatisticModelCode = "twoSampleTTest"
	// Used for paired t-test, may be paired with "value" to express degrees of freedom
	StatisticModelCodePairedTTest StatisticModelCode = "pairedTTest"
	// Used for Chi-square test, may be paired with "value" to express degrees of freedom
	StatisticModelCodeChiSquareTest StatisticModelCode = "chiSquareTest"
	// Used for Chi-square test for trend, may be paired with "value" to express degrees of freedom
	StatisticModelCodeChiSquareTestTrend StatisticModelCode = "chiSquareTestTrend"
	// Used for Pearson correlation, no additional elements needed
	StatisticModelCodePearsonCorrelation StatisticModelCode = "pearsonCorrelation"
	// Used for ANOVA method of analysis, may be paired with "value" to express degrees of freedom
	StatisticModelCodeAnova StatisticModelCode = "anova"
	// Used for one-way ANOVA method of analysis, may be paired with "value" to express degrees of freedom
	StatisticModelCodeAnovaOneWay StatisticModelCode = "anovaOneWay"
	// Used for 2-way ANOVA without replication method of analysis, may be paired with "value" to express degrees of freedom
	StatisticModelCodeAnovaTwoWay StatisticModelCode = "anovaTwoWay"
	// Used for 2-way ANOVA with replication method of analysis, may be paired with "value" to express degrees of freedom
	StatisticModelCodeAnovaTwoWayReplication StatisticModelCode = "anovaTwoWayReplication"
	// Used for multivariate ANOVA (MANOVA) method of analysis, may be paired with "value" to express degrees of freedom
	StatisticModelCodeManova StatisticModelCode = "manova"
	// Used for 3-way ANOVA method of analysis, may be paired with "value" to express degrees of freedom
	StatisticModelCodeAnovaThreeWay StatisticModelCode = "anovaThreeWay"
	// Used for sign test, no additional elements needed
	StatisticModelCodeSignTest StatisticModelCode = "signTest"
	// Used for Wilcoxon signed-rank test, no additional elements needed
	StatisticModelCodeWilcoxonSignedRankTest StatisticModelCode = "wilcoxonSignedRankTest"
	// Used for Wilcoxon rank-sum test, no additional elements needed
	StatisticModelCodeWilcoxonRankSumTest StatisticModelCode = "wilcoxonRankSumTest"
	// Used for Mann-Whitney U test, no additional elements needed
	StatisticModelCodeMannWhitneyUTest StatisticModelCode = "mannWhitneyUTest"
	// Used for Fisher's exact test, may be paired with "value" to express degrees of freedom
	StatisticModelCodeFishersExactTest StatisticModelCode = "fishersExactTest"
	// Used for McNemar's test, no additional elements needed
	StatisticModelCodeMcnemarsTest StatisticModelCode = "mcnemarsTest"
	// Used for Kruskal Wallis test, may be paired with "value" to express degrees of freedom
	StatisticModelCodeKruskalWallisTest StatisticModelCode = "kruskalWallisTest"
	// Used for Spearman correlation, no additional elements needed
	StatisticModelCodeSpearmanCorrelation StatisticModelCode = "spearmanCorrelation"
	// Used for Kendall correlation, no additional elements needed
	StatisticModelCodeKendallCorrelation StatisticModelCode = "kendallCorrelation"
	// Used for Friedman test, no additional elements needed
	StatisticModelCodeFriedmanTest StatisticModelCode = "friedmanTest"
	// Used for Goodman Kruska’s Gamma, no additional elements needed
	StatisticModelCodeGoodmanKruskasGamma StatisticModelCode = "goodmanKruskasGamma"
	// Used for GLM (Generalized Linear Model), no additional elements needed
	StatisticModelCodeGlm StatisticModelCode = "glm"
	// Used for GLM with probit link, no additional elements needed
	StatisticModelCodeGlmProbit StatisticModelCode = "glmProbit"
	// Used for GLM with logit link, no additional elements needed
	StatisticModelCodeGlmLogit StatisticModelCode = "glmLogit"
	// Used for GLM with identity link, no additional elements needed
	StatisticModelCodeGlmIdentity StatisticModelCode = "glmIdentity"
	// Used for GLM with log link, no additional elements needed
	StatisticModelCodeGlmLog StatisticModelCode = "glmLog"
	// Used for GLM with generalized logit link, no additional elements needed
	StatisticModelCodeGlmGeneralizedLogit StatisticModelCode = "glmGeneralizedLogit"
	// Used for Generalized linear mixed model (GLMM), no additional elements needed
	StatisticModelCodeGlmm StatisticModelCode = "glmm"
	// Used for GLMM with probit link, no additional elements needed
	StatisticModelCodeGlmmProbit StatisticModelCode = "glmmProbit"
	// Used for GLMM with logit link, no additional elements needed
	StatisticModelCodeGlmmLogit StatisticModelCode = "glmmLogit"
	// Used for GLMM with identity link, no additional elements needed
	StatisticModelCodeGlmmIdentity StatisticModelCode = "glmmIdentity"
	// Used for GLMM with log link, no additional elements needed
	StatisticModelCodeGlmmLog StatisticModelCode = "glmmLog"
	// Used for GLMM with generalized logit link, no additional elements needed
	StatisticModelCodeGlmmGeneralizedLogit StatisticModelCode = "glmmGeneralizedLogit"
	// Used for linear regression method of analysis, no additional elements needed
	StatisticModelCodeLinearRegression StatisticModelCode = "linearRegression"
	// Used for logistic regression method of analysis, no additional elements needed
	StatisticModelCodeLogisticRegression StatisticModelCode = "logisticRegression"
	// Used for Polynomial regression method of analysis, no additional elements needed
	StatisticModelCodePolynomialRegression StatisticModelCode = "polynomialRegression"
	// Used for Cox proportional hazards method of analysis, no additional elements needed
	StatisticModelCodeCoxProportionalHazards StatisticModelCode = "coxProportionalHazards"
	// Used for Binomial Distribution for Regression, no additional elements needed
	StatisticModelCodeBinomialDistributionRegression StatisticModelCode = "binomialDistributionRegression"
	// Used for Multinomial Distribution for Regression, no additional elements needed
	StatisticModelCodeMultinomialDistributionRegression StatisticModelCode = "multinomialDistributionRegression"
	// Used for Poisson Regression, no additional elements needed
	StatisticModelCodePoissonRegression StatisticModelCode = "poissonRegression"
	// Used for Negative Binomial Regression, no additional elements needed
	StatisticModelCodeNegativeBinomialRegression StatisticModelCode = "negativeBinomialRegression"
	// Zero-cell adjustment done by adding a constant to all cells of affected studies, paired with "value" to define the constant
	StatisticModelCodeZeroCellConstant StatisticModelCode = "zeroCellConstant"
	// Zero-cell adjustment done by treatment arm continuity correction, no additional elements needed
	StatisticModelCodeZeroCellContinuityCorrection StatisticModelCode = "zeroCellContinuityCorrection"
	// Used for adjusted analysis, paired with variable element(s)
	StatisticModelCodeAdjusted StatisticModelCode = "adjusted"
	// Used for interaction term, paired with "value" and two or more variable elements
	StatisticModelCodeInteractionTerm StatisticModelCode = "interactionTerm"
	// Used for Mantel-Haenszel method, no additional elements needed
	StatisticModelCodeManteHaenszelMethod StatisticModelCode = "manteHaenszelMethod"
	// Used for meta-analysis, no additional elements needed
	StatisticModelCodeMetaAnalysis StatisticModelCode = "metaAnalysis"
	// Used for inverse variance method of meta-analysis, no additional elements needed
	StatisticModelCodeInverseVariance StatisticModelCode = "inverseVariance"
	// Used for Peto method of meta-analysis, no additional elements needed
	StatisticModelCodePetoMethod StatisticModelCode = "petoMethod"
	// Hartung-Knapp/Hartung-Knapp-Sidik-Jonkman adjustment used in meta-analysis, no additional elements needed
	StatisticModelCodeHartungKnapp StatisticModelCode = "hartungKnapp"
	// Modified Hartung-Knapp/Hartung-Knapp-Sidik-Jonkman adjustment used in meta-analysis, no additional elements needed
	StatisticModelCodeModifiedHartungKnapp StatisticModelCode = "modifiedHartungKnapp"
	// From a fixed-effects analysis, no additional elements needed
	StatisticModelCodeEffectsFixed StatisticModelCode = "effectsFixed"
	// From a random-effects analysis, no additional elements needed
	StatisticModelCodeEffectsRandom StatisticModelCode = "effectsRandom"
	// Used for Chi-square test for homogeneity, may be paired with "value" to express degrees of freedom
	StatisticModelCodeChiSquareTestHomogeneity StatisticModelCode = "chiSquareTestHomogeneity"
	// Used for Dersimonian-Laird method of tau estimation, no additional elements needed
	StatisticModelCodeDersimonianLairdMethod StatisticModelCode = "dersimonianLairdMethod"
	// Used for Paule-Mandel method of tau estimation, no additional elements needed
	StatisticModelCodePauleMandelMethod StatisticModelCode = "pauleMandelMethod"
	// Used for Restricted Maximum Likelihood method of tau estimation, no additional elements needed
	StatisticModelCodeRestrictedLikelihood StatisticModelCode = "restrictedLikelihood"
	// Used for Maximum Likelihood method of tau estimation, no additional elements needed
	StatisticModelCodeMaximumLikelihood StatisticModelCode = "maximumLikelihood"
	// Used for Empirical Bayes method of tau estimation, no additional elements needed
	StatisticModelCodeEmpiricalBayes StatisticModelCode = "empiricalBayes"
	// Used for Hunter-Schmidt method of tau estimation, no additional elements needed
	StatisticModelCodeHunterSchmidt StatisticModelCode = "hunterSchmidt"
	// Used for Sidik-Jonkman method of tau estimation, no additional elements needed
	StatisticModelCodeSidikJonkman StatisticModelCode = "sidikJonkman"
	// Used for Hedges method of tau estimation, no additional elements needed
	StatisticModelCodeHedgesMethod StatisticModelCode = "hedgesMethod"
	// Dersimonian-Laird method for tau squared
	StatisticModelCodeTauDersimonianLaird StatisticModelCode = "tauDersimonianLaird"
	// Paule-Mandel method for tau squared
	StatisticModelCodeTauPauleMandel StatisticModelCode = "tauPauleMandel"
	// Restricted Maximum Likelihood method for tau squared
	StatisticModelCodeTauRestrictedMaximumLikelihood StatisticModelCode = "tauRestrictedMaximumLikelihood"
	// Maximum Likelihood method for tau squared
	StatisticModelCodeTauMaximumLikelihood StatisticModelCode = "tauMaximumLikelihood"
	// Empirical Bayes method for tau squared
	StatisticModelCodeTauEmpiricalBayes StatisticModelCode = "tauEmpiricalBayes"
	// Hunter-Schmidt method for tau squared
	StatisticModelCodeTauHunterSchmidt StatisticModelCode = "tauHunterSchmidt"
	// Sidik-Jonkman method for tau squared
	StatisticModelCodeTauSidikJonkman StatisticModelCode = "tauSidikJonkman"
	// Hedges method for tau squared
	StatisticModelCodeTauHedges StatisticModelCode = "tauHedges"
	// Mantel-Haenszel method for pooling in meta-analysis
	StatisticModelCodePoolMantelHaenzsel StatisticModelCode = "poolMantelHaenzsel"
	// Inverse variance method for pooling in meta-analysis
	StatisticModelCodePoolInverseVariance StatisticModelCode = "poolInverseVariance"
	// Peto method for pooling in meta-analysis
	StatisticModelCodePoolPeto StatisticModelCode = "poolPeto"
	// Generalized linear mixed model (GLMM) method for pooling in meta-analysis
	StatisticModelCodePoolGeneralizedLinearMixedModel StatisticModelCode = "poolGeneralizedLinearMixedModel"
)
