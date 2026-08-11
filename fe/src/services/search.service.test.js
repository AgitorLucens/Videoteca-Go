import { describe, it, expect, vi, beforeEach } from 'vitest'
import { searchMovies } from './search.service'
import { doGet } from './http.service'

vi.mock('./http.service', () => ({
  doGet: vi.fn(),
}))

describe('search.service', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  it('should call doGet with search query', async () => {
    const mockData = [{ id: 1, title: 'Inception' }]
    doGet.mockResolvedValue({ data: mockData })

    const result = await searchMovies('Inception')

    expect(doGet).toHaveBeenCalledWith('/search?q=Inception')
    expect(result).toEqual({ data: mockData })
  })

  it('should encode special characters in query', async () => {
    doGet.mockResolvedValue({ data: [] })

    await searchMovies('Test & Movie')

    expect(doGet).toHaveBeenCalledWith('/search?q=Test%20%26%20Movie')
  })

  it('should return empty array when no results', async () => {
    doGet.mockResolvedValue({ data: [] })

    const result = await searchMovies('xyz')

    expect(result.data).toEqual([])
  })
})